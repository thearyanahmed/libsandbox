# Kubernetes Cluster Architecture

**Master Node**

   - Communication using the API Server. Think of it as the frontend
   - Scheduling
   - Controller manager runs the controllers

K8s uses etcd for data store. Interact with **masternode** using kubectl. Master node manages worker node. Communicates with kubelet. **Kube-proxy** is the network routing for TCP and UDP. Docker daemon runs docker containers. Smallest unit is **pod**.

Worker nodes can be exposed to the internet via loadbalancers and traffic coming to the node is handled by kubeproxy.

**Nodes**

Node works as a worker machine. Node requirements:

   - Kubelete running
   - Supervisord
   - Kubeproxy
   - Container tooling

Recommended 3 node cluster.

**Pod**

Pod is one single unit of deployment. Most basic construct. Most simple unit. Create, deploy, delete pods and contains one running process. Pod Requirements:

   - Docker container
   - Storage resources
   - Unique network IP
   - Options that govern how container(s) should run

Pods don’t self heal. Never creates itself. Don’t use a pod directly, use controller.

Pod states:

   - Pending : Accepted by k8s but container not running
   - Running: At least one container running
   - Succeeded : All the containers successfully exited with exit status 0
   - Failed: All containers exited and at least 1 failed with non-zero exit status
   - CrashLoopBackOff: Self-explanatory

**Controllers**

Benefits of Controllers:

   - Application reliability, maybe a single pod is failing
   - Scaling
   - Load balancing

Kinds of controllers

   - **ReplicaSets**
      - Ensures a specificed number of replicas are running at all times. Needs to be used inside a deployment. Uses want vs have condition. New deployment will trigger new replica set while keeping the old ones, in case of rollbacks.
   - **Deployments**
      - Declartive updates config objects Allows us to do pod managment. Running replicate set allows us to deploy a number of pods and check their status as a single unit.
      - Sclaing a replicateSet sacles out the pods, allows for the deployment to handle more traffic.
      - Pause and resume During pause state, only updates are paused. Network will flow to existing ReplicaSets
      - Status
   - **DaemonSets**
      - Ensures all nodes run a copy of a specific pod. DaemonSets will add or remove the required pods.
   - **Jobs**
      - Supervisor process for pods carrying out batch jobs. Usually runs as a cron.
   - **Services**
      - Allows communication between one set of deployments with another. Solves the problem of new (dynamic) IP (I guess). Service as DNS.
      - Kinds:
         - Internal service: IP is reachable within the cluster
         - External : Endpoint available through node ip. NodePort.
         - Load Balancer: Exposes application to the internet with a load balancer.

**Labels, Selectors and Namespaces**

   Labes:

      - Key-Value pairs. For developers, image tags with value. Label keys are unique per object.

   Selectors:

      - Equality based selectors. `=` and `!=` .
      - Set based selectors. Has: `IN` , `NOTIN` and `EXISTS`

   Namespaces:

      - Multiple virtual clusters backed by a single physical server. Mainly for separation or concern and divide cluster resoruces.
      - K8s starts with default namespace. Newer applications install with their resources in a different namespace.

**Kubelet**

   - Communicates with API servers for pods
   - Executes pod container via container engine
   - Mounts and runs pod vols and secrets
   - Executes health checks

Podspec: YAML file that describes a pod. Uses API server.

**Kube-proxy**

Runs on all the worker nodes. Reflects services as defined on each node. Can do simple network stream or round-robin forwarding.

**Modes**

   - User space mode
   - Iptables mode
   - Ipvs mode ( alpha feature

For a new service, kube-proxy opens a randomly chosen port on the local node.

**What does ClusterIP, NodePort, and LoadBalancer mean?**

- **ClusterIP**
   - The default value. The service is only accessible from within the Kubernetes cluster. You can’t make requests to your Pods from outside the cluster!
- **NodePort**
   - This makes the service accessible on a static port on each Node in the cluster. This means that the service can handle requests that originate from outside the cluster.
- **LoadBalancer**
   - The service becomes accessible externally through a cloud provider's load balancer functionality. GCP, AWS, Azure, and OpenStack offer this functionality. The cloud provider will create a load balancer, which then automatically routes requests to your Kubernetes Service

**Service config**

```go
kind: Service 
apiVersion: v1 
metadata:
  name: hostname-service 
spec:
  # Expose the service on a static port on each node
  # so that we can access the service from outside the cluster 
  type: NodePort

  # When the node receives a request on the static port (30163)
  # "select pods with the label 'app' set to 'echo-hostname'"
  # select pod where label = 'app' and label_value = 'echo-hostname'
  # and forward the request to one of them
  selector:
    app: echo-hostname 

  ports:
    # Three types of ports for a service
    # nodePort - a static port assigned on each the node
    # port - port exposed internally in the cluster
    # targetPort - the container port to send requests to
    - nodePort: 30163
      port: 8080 // the service is internally available on this port
      targetPort: 80 // the container running inside the port exposes this port
```

**Kubectl Commands**

- kubectl get pods
- kubectl get services
- kubectl create -f $filename.yml
- kubectl get replicas
- kubectl get pods —show-labels
- kubectl get pods —selector dev-lead=$name
- kubectl get pods —selectors dev-lead=$name,env=staging
- kubectl get pods —selectors dev-lead!=$name # not equals operator
- kubectl get pods -l # shortcuts for selectors
- kubectl get pods -l ‘release version in (1.0, 2.0)’
- kubectl get pods -l ‘release version notin (1.0, 2.0)’  # not in operator
- kubectl delete pods -l deve-lead=$name
- kubectl create -f $config.yml —record # to record rollout history
- kubectl set image deployment/$deploymentName helloworld=$image:$tag # updates image
- kubectl rollout history deployment/$name_of_deployment
- kubectl rollout undor deployment/$name_of_deployment
- kubectl rollout undor deployment/$name_of_deployment —to-revision=$version_number
- kubectl describe deployment $deployment_name
- kubectl describe $type $name
- kubectl logs $pod_name
- kubectl exec -it $pod_name $command_to_run # ($command_to_run eg: /bin/bash)
- kubectl autoscale deployment php-apache --cpu-percent=50 --min=1 --max=10
- kubectl get hpa
- kubectl create namespace production
- kubectl describe pod nginx -n development
- kubectl delete pod nginx -n development

**Minikube commands**

- minikube service  $service-name

**Liveness probe vs Readiness probe**

Readiness states if a container (? pod?)  is **ready** to serve requests or not. Liveness probe states if the container is running or not. Maybe it started successfully and then crashed for some reason.

