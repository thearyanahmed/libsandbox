from random import randrange
import os
from bootstrap import load_env
from redis_cluster_client import RedisClusterClient

load_env()

total = 0

data = []

with open('data.txt', 'r') as f:
    for x in f:
        total += 1
        data.append(x.rstrip())

redis_cluster = RedisClusterClient()

print("Reading from {0} hashes".format(total))

for x in range(int(os.getenv("READ_COUNT"))):
    index = randrange(0, total)
    key = data[index]
    print("{0} => {1}".format(key,redis_cluster.get(key)))
