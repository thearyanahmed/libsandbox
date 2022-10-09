# 03 | Producer

Producers write messages to Kafka that can then be consumed by readers.

### Write workflow

![Image.png](https://res.craft.do/user/full/ff254a62-8b8f-1fb3-8fb0-7eba2a9eb4d4/doc/F290208A-B03E-4CDC-BA7F-3F7113B5CD16/478C3DEB-208B-427B-9419-6CCB4358EBE3_2/ZB1RRvXxyqxiTXKO6FlidtkSVgHNfiDIR2xC1jyYmR8z/Image.png)

1. An object of class `ProduceRecord` is instantiated which must contain the intended topic of the message and the message itself, which is the value. The message key and partition can optionally be included.
2. Next, the key and values to be sent over the network are serialized. At this point, a serialization exception can be thrown if serialization fails. Two properties `key.serializer` and `value.serializer` need to be set to the classes that know how to convert the key and value objects to bytes.
3. The data is next sent to the partitioner. If the partition has been specified in `ProduceRecord` then the partitioner takes no action and simply returns the already specified partition. If the partition is not specified, the partitioner selects a partition based on the key.
4. Once the partition has been determined, the producer adds the message to a batch of records waiting to be sent to the same topic and partition. A different thread is responsible for sending the batches of records to the right Kafka brokers. The location of the Kafka cluster is gleaned from the property `bootstrap.servers`, which is a list of `host:port` pairs that the producer will attempt to connect to. This doesn’t need an exhaustive list of brokers since the producer receives additional information after the initial connection.
5. Once the broker receives the batch of records and writes the messages to Kafka successfully, it returns an object of `RecordMetadata` that contains information about the topic, partition, and offset of the record within the partition.
6. If the broker fails to write the messages, it will return an error to the producer. The producer attempts a few retries before giving up.

