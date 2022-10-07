# Redis clustering 

Redis Cluster provides a way to run a Redis installation where data is  **automatically sharded across multiple Redis nodes**.

## Table of content
- [What is redis](#what-is-redis)
- [Redis Cluster](#redis-cluster)
- [Redis Cluster Goals](#redis-cluster-goals)
- [Sharding](#sharding)
- [Detecting failures](#detecting-failures)
- [Clustering in your own machine](#a-cluster-in-your-own-machine)
- [Simulating-failover](#simulating-failover)
- [Running this project](#how-to-run-the-project)
- [Notes](#notes-to-self)

### What is Redis?

1.  Redis is an  **in-memory,**  **key-value**  **store**.
-   **In-memory store**: Redis keeps the data in the cache and it does not write to the disk. This makes reading/writing data very fast. (However, Redis has an option to write data to the disk)
-   **Key-value store**: Redis can store data as key-value pairs.
2. It is a  [No-SQL](https://en.wikipedia.org/wiki/NoSQL)  database.
3. Uses  [data structures](https://redis.io/topics/data-types-intro)  to store data.
4. Interaction with data is command-based.
 
 for example
```bash
SET $KEY $VALUE
#eg:
SET name Aryan

#if your key or value has spaces, wrap them with quotes
SET $KEY "$VALUE" 
SET 'full-name' 'Aryan Ahmed Anik'
```

**We can put the requests that need to be serviced very fast into the Redis memory and service from there while keeping the rest of the data in the main database as it is a in-memeory database.**

### Redis cluster
<p align="center">
  <img width="460" height="300" src="https://i.ibb.co/KWgJBMF/redis-cluster-diagram.png">
</p>

A Redis cluster is simply a **data sharding strategy**. 

### Sharding
Redis Cluster does not use consistent hashing, but a different form of sharding where every key is conceptually part of what we call an  **hash slot**.

There are 16384 hash slots in Redis Cluster, and to compute what is the hash slot of a given key, we simply take the CRC16 of the key modulo 16384.

Every node in a Redis Cluster is responsible for a subset of the hash slots, so for example you may have a cluster with 3 nodes, where:

-   Node A contains hash slots from 0 to 5500.
-   Node B contains hash slots from 5501 to 11000.
-   Node C contains hash slots from 11001 to 16383.

#### What is sharding?
  
The word “**Shard**” means “**a small part of a whole**“. Hence Sharding means dividing a larger part into smaller parts.
Sharding is a type of partitioning in which a large *data store* is divided or partitioned into smaller data, also known as shards.
Redis cluster  automatically partitions data across multiple Redis nodes. It is an advanced version of Redis that achieves distributed storage and prevents a single point of failure.
 
 -   Horizontally scalable: We can continue to add more nodes as the capacity requirement increases.
-   Auto data sharding: can automatically partition and split data among the nodes.
-   Fault tolerant: even though we lose a node, we can still continue to operate without losing any data and no downtime.
-   Decentralized cluster management: no single node acts as an orchestrator of the entire cluster, every node participates in the cluster configuration ([via gossip protocol](https://en.wikipedia.org/wiki/Gossip_protocol)).

So in summary, 
- The ability to  **automatically split your dataset among multiple nodes**.
- The ability to  **continue operations when a subset of the nodes are experiencing failures**  or are unable to communicate with the rest of the cluster.
-  *It also provides **some degree of availability during partitions**, that is in practical terms the ability to continue the operations when some nodes fail or are not able to communicate.*

 
### Redis Cluster goals

Redis Cluster is a distributed implementation of Redis with the following goals, in order of importance in the design:

-   High performance and linear scalability up to 1000 nodes. There are no proxies, asynchronous replication is used, and no merge operations are performed on values.
-   Acceptable degree of write safety: the system tries (in a best-effort way) to retain all the writes originating from clients connected with the majority of the master nodes. Usually there are small windows where acknowledged writes can be lost. Windows to lose acknowledged writes are larger when clients are in a minority partition.
-   Availability: Redis Cluster is able to survive partitions where the majority of the master nodes are reachable and there is at least one reachable slave for every master node that is no longer reachable. Moreover using  _replicas migration_, masters no longer replicated by any slave will receive one from a master which is covered by multiple slaves.

  
Every Redis Cluster node requires two TCP connections open. The normal Redis TCP port used to serve clients, for example 6379, plus the port obtained by adding 10000 to the data port, so 16379 in the example.
**The offset is always 10000**.

### Prerequisites

The minimal cluster that works as expected requires;
-   Minimum 3 Redis master nodes
-   Minimum 3 Redis slaves, 1 slave per master (to allow minimal fail-over mechanism)

### Distributed storage of the cluster

Every key that you save into a Redis cluster is associated with a hash slot. There are 0–16383 slots in a Redis cluster. Thus, a Redis cluster can have a maximum of 16384 master nodes (with the first 0) and each master node in a cluster handles a subset of the 16384 hash slots.

### Cyclic Redundency Check

A **cyclic redundancy check** (**CRC**) is an [error-detecting code](https://en.wikipedia.org/wiki/Error_detection_and_correction "Error detection and correction") commonly used in digital [networks](https://en.wikipedia.org/wiki/Telecommunications_network "Telecommunications network") and storage devices to detect accidental changes to raw data. Blocks of data entering these systems get a short _check value_ attached, based on the remainder of a [polynomial division](https://en.wikipedia.org/wiki/Polynomial_long_division "Polynomial long division") of their contents. On retrieval, the calculation is repeated and, in the event the check values do not match, corrective action can be taken against data corruption. CRCs can be used for [error correction](https://en.wikipedia.org/wiki/Error_correcting_code "Error correcting code") (see [bitfilters](https://en.wikipedia.org/wiki/Mathematics_of_cyclic_redundancy_checks#Bitfilters "Mathematics of cyclic redundancy checks")).[[1]](https://en.wikipedia.org/wiki/Cyclic_redundancy_check#cite_note-1)
-wikipedia

```
HASH_SLOT = CRC16(key) mod HASH_SLOTS_NUMBER
```


For example, let’s assume the key space is divided into 10 slots (0–9). Each node will hold a subset of the hash slots.
<p align="center">
  <img src="https://miro.medium.com/max/1400/1*g_uPH1TnC40Nqiqx4X0ifQ.png">
</p>

### Detecting failures

Every node has a unique ID in the cluster. This ID is used to identify each and every node across the entire cluster using the gossip protocol.

So, a node keeps the following information within itself;

-   Node ID, IP, and port
-   A set of flags
-   What is the Master of the node if it is flagged as “slave”
-   Last time a node was pinged
-   Last time the pong was received

Nodes in a cluster always keep gossiping with each other, enabling them to auto-discover other nodes.

e.g. If A knows B, and B knows C, eventually B will send gossip messages to A about C. Then A will register C as part of the network and will try to connect with C.

When node A talks to the cluster by the gossip protocol, If a/any node fails, the cluster's reply will be node X,Y responded okay but P,Q,R didn't response.

### A cluster in your own machine
Let's work through by building a cluster ourself. At first we need configs for our cluster. Remember, a redis cluster is like a simple redis instace but running inside a cluster. You can use the default config that comes with redis ( $your_redis_instanlation_dir / redis.conf)  and enable clustering.

Let say we'll have 3 master and 3 salves appointed with each paster, our ports will start from 7000 to 7005. 
So, lets make a directory `cluster` and inside it make configuration file for each node, you can name them what ever you want, I'm calling them by their port.conf, so 7000.conf, 7001.conf .... upto 7005.conf . 

Add the following to each of the config files, change the port, cluster-node-*.conf, appendfile name, dbfile name respectively with their port (can be different though, but I'm trying to keep them consistant and its easier to read/understand).  

```bash
port 7000
cluster-enabled yes
cluster-config-file cluster-node-0.conf
cluster-node-timeout 5000
appendonly yes
appendfilename node-0.aof
dbfilename dump-0.rdb
``` 

### Explaining the config file
These are *redis cluster configuration parameters*.

-   **cluster-enabled  `<yes/no>`**: If yes, enables Redis Cluster support in a specific Redis instance. Otherwise the instance starts as a stand alone instance as usual.
-   **cluster-config-file  `<filename>`**: Note that despite the name of this option, this is not a user editable configuration file, but the file where a Redis Cluster node automatically persists the cluster configuration (the state, basically) every time there is a change, in order to be able to re-read it at startup. The file lists things like the other nodes in the cluster, their state, persistent variables, and so forth. Often this file is rewritten and flushed on disk as a result of some message reception.
-   **cluster-node-timeout  `<milliseconds>`**: The maximum amount of time a Redis Cluster node can be unavailable, without it being considered as failing. If a master node is not reachable for more than the specified amount of time, it will be failed over by its slaves. This parameter controls other important things in Redis Cluster. Notably, every node that can't reach the majority of master nodes for the specified amount of time, will stop accepting queries.
-   **cluster-slave-validity-factor  `<factor>`**: If set to zero, a slave will always try to failover a master, regardless of the amount of time the link between the master and the slave remained disconnected. If the value is positive, a maximum disconnection time is calculated as the  _node timeout_  value multiplied by the factor provided with this option, and if the node is a slave, it will not try to start a failover if the master link was disconnected for more than the specified amount of time. For example if the node timeout is set to 5 seconds, and the validity factor is set to 10, a slave disconnected from the master for more than 50 seconds will not try to failover its master. Note that any value different than zero may result in Redis Cluster to be unavailable after a master failure if there is no slave able to failover it. In that case the cluster will return back available only when the original master rejoins the cluster.
-   **cluster-migration-barrier  `<count>`**: Minimum number of slaves a master will remain connected with, for another slave to migrate to a master which is no longer covered by any slave. See the appropriate section about replica migration in this tutorial for more information.
-   **cluster-require-full-coverage  `<yes/no>`**: If this is set to yes, as it is by default, the cluster stops accepting writes if some percentage of the key space is not covered by any node. If the option is set to no, the cluster will still serve queries even if only requests about a subset of keys can be processed.
-   **cluster-allow-reads-when-down  `<yes/no>`**: If this is set to no, as it is by default, a node in a Redis Cluster will stop serving all traffic when the cluster is marked as fail, either when a node can't reach a quorum of masters or full coverage is not met. This prevents reading potentially inconsistent data from a node that is unaware of changes in the cluster. This option can be set to yes to allow reads from a node during the fail state, which is useful for applications that want to prioritize read availability but still want to prevent inconsistent writes. It can also be used for when using Redis Cluster with only one or two shards, as it allows the nodes to continue serving writes when a master fails but automatic failover is impossible.

We also have 

- **appendfilename** for AOF name.  
- **dbfilename** for data store. 

Now, we need to run each of them as redis-server instance with our own .conf file. 
```bash
prophecy$ redis-server cluster/7000.conf
prophecy$ redis-server cluster/7001.conf
...
prophecy$ redis-server cluster/N.conf
```
![redis-cluster](https://i.ibb.co/7KxLcs9/image.png)
After that,  we need to create our cluster with these redis-server instances.
```bash
redis-cli --cluster create 127.0.0.1:7000 127.0.0.1:7001 \
127.0.0.1:7002 127.0.0.1:7003 127.0.0.1:7004 127.0.0.1:7005 \
--cluster-replicas 1
```
If everything is okay, we should receive a response like
```bash
[OK] All 16384 slots covered
```

![enter image description here](https://i.ibb.co/KD9d5nv/image.png)
If we check the individual instances, we will see them initilize themselves as master or slaves, alocating hash slots etc.
After successsfully creating the cluster, our instances should be like 
![enter image description here](https://i.ibb.co/r6xYp4g/image.png)
Lets test our client,

```
# -c for cluster mode. 
redis-cli -c -p 7000
SET prophecy thearyanahmed
# output is 'OK'

# ^ + c to quit 
# login to a different node 

redis-cli -c -p 7005 
GET prophecy
# output is 
Redirecting to slot [$a_slot_number] $host:$port
"thearyanahmed"
```
In my case

![enter image description here](https://i.ibb.co/hVhkr3m/image.png)

### Simulating failover

So despite us setting our key value in 7000 node, we can reach it from 7005 (a different node).
Let's set some value in in 7001 node and then kill it. 
Now if we try to get that value, we can still get the follow, check the sequence of commands and results.

![redis cluster failover](https://i.ibb.co/fkRdB1n/image.png)

Why did that happen? Well, at first, 7001 was our master, but then, when it failed, the salve of 7001 became the master, and as data was replicated, every request was seem-lessly servered by the salve ( acting as the master node).

It self is not magic, it's pure logic, acting like magic! 

### How to run the project

We have a simple python project to connect to our cluster, wrting and reading out data from the nodes.
at first run
```bash
git clone https://github.com/thearyanahmed/redis-cluster.git redis-cluster
```
#### Project details

We have primarily 3 files,
- redis_cluster_client.py
- write_to_cluster.py
- read_from_cluster.py 

`redis_cluster_client` is the bridge for connecting  to the redis cluster we've created. `write_to_cluster` is used write some random data in the cluster. `read_from_cluster` is used to read data from the cluster randomly. 
We have `.env` file to setup some paramters, at first , make a copy of .env file from .env.example. 

#### .env params

- **WORKER_COUNT** is an integer value, pointing how many workers will be inserting data.
- **SINGLE_WORKER_TASK_COUNT** is an integer, it denotes number of data will be written by a single worker.
- **READ_COUNT** is an integer, when we use `read_from_cluster`, how many data are we attempting to read.
- **CLUSTER_HOST** is simply the host ip of your cluster.
- **CLUSTER_PORT_RANGE_STARTS** is an integer, denoting the starting port number.
- **CLUSTER_PORT_RANGE_ENDS** is an integer, denoting the ending port number.

If we have 3 in WORKER COUNT and 10 in SINGLE_WORKER_COUNT, total 3 x 10 = 30 data will be written.  

The application will use the CLUSTER_PORT_RANGE_* to generate port numbers to connect to the cluster. For example, if you have 50 nodes, form 7001 to 7050. No need to write 7001,7002,7003, ... , 7050. Just write `CLUSTER_PORT_RANGE_STARTS=7001` and `CLUSTER_PORT_RANGE_ENDS=7050`. 

The application needs python version 3, [python-dotenv](https://pypi.org/project/python-dotenv/), [redis-py-cluster](https://github.com/Grokzen/redis-py-cluster)  and [faker](https://faker.readthedocs.io/en/master/) . 

The application will run asyncronously for writing to the cluster.
To write, 
```bash
python3 write_to_cluster.py
```
To read, 
```bash
python3 read_from_cluster.py
```
	
### Notes to self

#### Replication Or Clustering?
- If you have more data than RAM in a single machine, use a Redis cluster to shard the data across multiple databases.

- If you have less data than RAM in a machine, set up a master/slave replication with a sentinel in front to handle the failover.

#### Why Do you need a minimum of 3 masters?

During the failure detection, the majority of the master nodes are required to come to an agreement. If there are only 2 masters, say A and B and B failed, then the A master node cannot reach to a decision according to the protocol. The A node needs another third node, say C, to tell A that it also cannot reach B.
You can have odd number of master nodes so the response if a node has failed or not, the response can't be even, like if you have 6, 3 might say it is still alive and 3 might say it is dead. Though it is highly unlikely to happen. But theoretically it can happen.

Special thanks to [Varuni Punchihewa](https://blog.usejournal.com/@varunipunchihewa) 
