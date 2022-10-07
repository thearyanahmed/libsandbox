from bootstrap import load_env
from rediscluster import RedisCluster
import os


class RedisClusterClient:
    """
    This is a cluster client. A wrapper over cluster-py library.
    """

    def __init__(self):
        load_env()
        startup_nodes = self.__get_cluster_nodes()

        # Note: decode_responses must be set to True when used with python3
        redis_cluster = RedisCluster(startup_nodes=startup_nodes, decode_responses=True)

        if redis_cluster is None:
            print("Could not connect to redis cluster.Exiting.")
            exit(1)

        self.__redis_cluster = redis_cluster

    def __get_cluster_nodes(self):
        nodes = []
        start = int(os.getenv("CLUSTER_PORT_RANGE_STARTS"))
        end = int(os.getenv("CLUSTER_PORT_RANGE_ENDS"))
        host = os.getenv("CLUSTER_HOST")

        for port in range(start, end + 1):
            nodes.append({"host": host, "port": port})
        return nodes

    # returns the redis_cluster library instance
    def instance(self):
        return self.__redis_cluster

    # writes a value to the cluster
    def set(self, key, value, ex, px, nx, xx):
        return self.__redis_cluster.set(key, value, ex, px, nx, xx)

    # reads from cluster
    def get(self, key):
        return self.__redis_cluster.get(key)
