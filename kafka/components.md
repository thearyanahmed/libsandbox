# 02 | Components of Kafka

This lesson explains the various participants in the Kafka ecosystem.

### Message

A message is simply an array of bytes from Kafka’s perspective. A message is also the unit of data in the Kafka ecosystem. Similar to row in RDMS.

- Messages are batched rather than being sent individually to reduce overhead. ::Throughput gets priority::. This leads to the classical tradeoff between latency and throughput. As batch sizes grow larger, throughput increases as more messages are handled per unit of time. At the same time, it takes longer for an individual message to be delivered thus increasing latency.
- Messages batched together are compressed for efficient data transfer.

### Topic

Messages get written and read from topics. A topic can be thought of as analogous to a directory or folder in an operating system.

### Partition

A topic has subdivisions known as partitions. There are a number of salient points to remember about a topic:

- Messages are ordered by time only within a partition and not across the entire topic.
- Messages are read from beginning to end in a partition.
- Message can only be appended to the end of a partition. (Notice the similarity to a commit log)
- Partitions allow Kafka to scale horizontally and also provide redundancy. Each partition can be hosted on a different server, which allows new partitions to be added to a topic as the load on the system increases.

![Image.png](https://res.craft.do/user/full/ff254a62-8b8f-1fb3-8fb0-7eba2a9eb4d4/doc/84BABF94-CC65-4014-8FB0-3609FCFBBFF5/555FE555-9351-4999-A11B-A75E6E8FC6E1_2/cOjeoftFwuGsPO4GsNSGhifYvNCtLEyC9DBihD3by9Uz/Image.png)

### Message key

All message with the same key are written in the same partition.

### Message offset

Messages also have a metadata associated with them called the **offset**. The offset is an integer value that is ever increasing and determines the order of the message within a partition. By remembering the message offset, the consumer is able to continue from where it previously left off.

### Brokers

A single Kafka server is called a **broker**. Usually, several Kafka brokers operate as one Kafka cluster. The cluster is controlled by one of the brokers, called the controller, which is responsible for administrative actions such as assigning partitions to other brokers and monitoring for failures. The controller is elected from the live members of the cluster.

Within a cluster, a single broker owns a partition and is said to be the **leader**. All the other partition-replicating brokers are called **followers**.

### Producers

Producers create messages and are sometimes known as writers or publishers. Producers can direct messages to specific partitions using the message key and implement complex rules for partition assignment using a custom partitioner.

### Consumers

Consumers read messages and are sometimes known as subscribers or readers. Consumers operate as a group called the **consumer group**. A consumer group consists of several consumers working together to read a topic.

It is also possible for a consumer group to have a single consumer in it. Each partition is read by a single member of the group, though a single consumer can read multiple partitions. The mapping of a consumer to a partition is called the ownership of the partition by the consumer.

If a consumer fails, the remaining consumers in the group will rebalance the partitions amongst themselves to make up for the failed member.

![Image.png](https://res.craft.do/user/full/ff254a62-8b8f-1fb3-8fb0-7eba2a9eb4d4/doc/84BABF94-CC65-4014-8FB0-3609FCFBBFF5/80697201-EBE0-41BE-8B1E-250FF5AA0658_2/eHUNLyQXmvCRjcjIwL5XBZ9qlexjKr7yMHr1jknCTuwz/Image.png)

