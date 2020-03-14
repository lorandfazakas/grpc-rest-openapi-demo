# grpc-rest-openapi-demo
This is a demo project for my [blog](https://lorand.dev) showcasing GRPC, REST and OpenApi.

## Services
* List all the ToDos
* Create a new ToDo
* Get a ToDo by ID
* Mark a ToDo as done


## How to use it

* You can start the server with the following command:

```cmd
go run server/server.go
```

* You can start the demo client which will get all the ToDos using gRPC:
```cmd
go run client/client.go
```

You can access the swagger-ui on http://localhost:10001/swaggerui/