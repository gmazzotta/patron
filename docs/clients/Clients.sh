#!/bin/sh
cat << EOF
# Clients

Patron microservices can interact with other microservices, APIs and applications using a number of clients.

All clients contain integrated tracing powered by \`opentracing-go\`; any new clients should attempt to do the same.

**Third-party dependencies**  
$(cat ../../go.mod | grep github.com/opentracing/opentracing-go | xargs)


## HTTP Client
Patron provides an HTTP client which integrates tracing into all outgoing requests by wrapping the default \`net/http\` client. Users can configure the client's Timeout, RoundTripper and/or set up a circuit breaker.


## AMQP
The AMQP client allows users to connect to a RabbitMQ instance and publish JSON or Protobuf messages using a fan-out exchange. The published messages have integrated tracing headers by default. Users can configure the timeout for connecting to the RabbitMQ instance.

**Third-party dependencies**  
$(cat ../../go.mod | grep github.com/streadway/amqp | xargs)

## GRPC
The GRPC client initiates a client connection to a given target while injecting a \`UnaryInterceptor\` to integrate tracing capabilities. By default, this is a non-blocking connection and users can pass in any number of [\`grpc.DialOption\`](https://github.com/grpc/grpc-go/blob/master/dialoptions.go) arguments to configure its behavior.

**Third-party dependencies**  
$(cat ../../go.mod | grep google.golang.org/grpc | xargs)


## Kafka
The Kafka client allows users to create a synchronous or asynchronous Kafka producer and publish Kafka messages with tracing headers. The builder pattern allows users to configure the Kafka version, the dial timeout, the RequiredAcks policy, as well as the encoder used to serialize the messages.

**Third-party dependencies**  
$(cat ../../go.mod | grep github.com/Shopify/sarama | xargs)


## Redis
The Redis client allows users to connect to a Redis instance and execute commands. The connection can be configured using [\`redis.Options\`](https://github.com/go-redis/redis/blob/v7/options.go).

**Third-party dependencies**  
$(cat ../../go.mod | grep github.com/go-redis/redis/v7 | xargs)


## SQL
The SQL client enhances the standard library SQL by integrating tracing capabilities. It has support for prepared statements, queries, as well as low-level handling of transactions.


## SNS - SQS
The SNS and SQS clients provide a set of interfaces and structs useful for publishing messages to AWS SNS and SQS, with integrating tracing.

**Third-party dependencies**  
$(cat ../../go.mod | grep github.com/aws/aws-sdk-go | xargs)


## Elasticsearch
The Elasticsearch client allows users to connect to an elasticsearch instance. Its behavior can be configured by providing an [\`elasticsearch.Config\`](https://github.com/elastic/go-elasticsearch/blob/4b40206692088570801280584e614027e6ce818b/elasticsearch.go#L32) struct

**Third-party dependencies**  
$(cat ../../go.mod | grep github.com/elastic/go-elasticsearch | xargs)

EOF