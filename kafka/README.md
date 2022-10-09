# Building scalable data pipelines with Kafka

### TOC
 - Introduction (current page)
 - [Components](components.md)
 - [Producers](producers.md)


### What is Kafka?

> A distributed event streaming platform that lets you read, write, store, and process events (also called records or messages in the documentation) across many machines.

Commands for startup, create topics, produce and consume messages

```Bash
# Change directory to Kafka's installation on directory
cd /Kafka

# Run Zookeeper from the bin folder but redirect the outputs to null so that
# messages from Zookeeper don't get mixed-up with messages from Kafka on the console
 bin/zookeeper-server-start.sh config/zookeeper.properties > /dev/null 2>&1 &

 # You can run the following command to verify that Zookeeper is indeed running
 ps -aef

 # Run the  Kafka service. You should see output from the service on the console. 
 # Hit Enter key when the output stops to return to the console.
 bin/kafka-server-start.sh config/server.properties &

 # At this point you have a basic Kafka environment setup and running.

 # Create a topic 'datajek-events' to publish to.
 bin/kafka-topics.sh --create --topic datajek-topic --bootstrap-server localhost:9092

# Write some messages to the 'datajek-events' topic
bin/kafka-console-producer.sh --topic datajek-topic --bootstrap-server localhost:9092

# Press Ctrl+C to stop the producer

# Read the messages written to the topic by running the consumer
bin/kafka-console-consumer.sh --topic datajek-topic --from-beginning --bootstrap-server localhost:9092

# You should see all the messages you typed earlier, displayed on the console. Press
# Ctrl+C anytime to stop the consumer.
```

### Characteristics of Distributed SystemsReliability is a system’s ability to continue to work correctly in spite of faults. A reliable system can be banked upon to continue to work without degradation of service if a part of the overall system fails.

Reliability is a system’s ability to continue to work correctly in spite of faults. A reliable system can be banked upon to continue to work without degradation of service if a part of the overall system fails.

- Reliability
- Fault vs. Failure
   - When a part of a system experiences failure but the system as a whole can still continue to operate reliably, we label the system as **fault-tolerant** or **resilient.**
- Scalability
   - Horizontal : Multiple servers.
   - Vertical : Same server, upgrade.
- Measure performance
- Maintainability
- Availability
- Consistency

### Message patterns

- **Pub - Sub**
   - In the publish-subscribe model, a participant in the system produces data and publishes the data to a channel or topic. The message can be consumed by multiple readers and the messages are always delivered to the consumers in the order that they were published in.

![Image.png](https://res.craft.do/user/full/ff254a62-8b8f-1fb3-8fb0-7eba2a9eb4d4/doc/4E0BDFCA-B578-4F21-B938-902F3057DC70/A5147DEC-08A0-473D-AA08-4CBAFBECC7FB_2/IracdgVW9Qh3mZPvI0nCuWwgfvBdowYOX0JwXGWGOo0z/Image.png)

- **Message Queues**
   - In contrast to pub-sub pattern, message queuing publishes a message to a topic or channel which is processed exactly once by one consumer. Once the message has been processed and the consumer acknowledges ::consumption of the message, the message is deleted:: from the queue. Implementation dictates which consumer a message will be delivered to for processing.

![Image.png](https://res.craft.do/user/full/ff254a62-8b8f-1fb3-8fb0-7eba2a9eb4d4/doc/4E0BDFCA-B578-4F21-B938-902F3057DC70/C624C598-C77F-466C-9E77-EF24DFC1A279_2/4y9bJawJqtZvgUHupppHULLzsWc8HoWyeAPSgYMnvWYz/Image.png)

