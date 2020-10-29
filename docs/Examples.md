# Examples

The examples have been created in order to be able to start them all and see how a request travels throughout all services.
Detailed examples can be found in the [examples](/examples) folder.

## Prerequisites

Before all services start we should first start all dependencies with `docker-compose`.

```shell
docker-compose up -d
```

To tear down the above just:

```shell
docker-compose down
```

## [HTTP Cache Service](../examples/http-cache/main.go) 

The service shows how to use:
 
- HTTP caching in a specific route using Redis
- Textual logger 
  
The service can be started as follows:

```shell
go run examples/http-cache/main.go
```

## [HTTP Service](../examples/http/main.go)

The service shows how to use:
 
- HTTP with CORS
- HTTP client with API key
- Textual logger with predefined fields
  
The service can be started as follows:

```shell
go run examples/http/main.go
```

## [HTTP API Key Service](../examples/http-sec/main.go)

- HTTP service with a secured route (API KEY)
- Async Kafka publisher
- Default structured logger with predefined fields

The service can be started as follows:

```shell
go run examples/http-sec/main.go
```

## [Kafka Service](../examples/kafka/main.go)

The service shows how to use:

- Kafka with a group consumer
- AMQP publisher
- Textual logger

The service can be started as follows:

```shell
go run examples/amqp/main.go
```

## [AMQP Service](../examples/kafka/main.go)

The service shows how to use:

- AMQP consumer
- AWS SNS Publisher
- AWS SQS Publisher
- Default structured logger

The service can be started as follows:

```shell
go run examples/sns/main.go
```

## [AWS SQS Service](../examples/amqp/main.go)

The service shows how to use:

- AWS SQS Consumer
- gRPC client
- Default structured logger

The service can be started as follows:

```shell
go run examples/sqs/main.go
```

## [gRPC Service](../examples/grpc/main.go)

The service shows how to use:

- gRPC Server
- Textual logger

The service can be started as follows:

```shell
go run examples/grpc/main.go
```

## All of them working together

After all service have been started successfully we can send a request,  
and see how it travels through all of our services.
and then send a sample request:

```shell
../examples/start_processing.sh
```

After that head over to [jaeger](http://localhost:16686/search) and [prometheus](http://localhost:9090/graph).