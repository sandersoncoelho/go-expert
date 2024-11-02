# Clean-Architecture

In this challenge it was implemented the orders listing for rest, grpc and graphql methods. For that we used the basic structure available at https://github.com/devfullcycle/goexpert/tree/main/20-CleanArch.

### Database

It is available a docker-compose.yaml to start a MySQL container to persist our application data. At the root diretory application /go-expert/clean-architecture, run:

```
$ docker compose up -d
```

In our Makefile, there are commands to migrate our structure table to MySQL, but for create our tables, we just need to run:

```
$ make migrate
```

### Warnings

Before run the application, check if the dependencies are resolved, if not, do `go mod tidy`. Check also if the Evans client is available, otherwise do `PATH="$PATH:$(go env GOPATH)/bin"`. And check if the ports setted at cmd/ordersystem/.env is available.

```
sudo lsof -i:8000
sudo lsof -i:50051
sudo lsof -i:8080
```

If these commands returned anything, it means the ports are available to run our application.

### Running the application

```
$ cd cmd/ordersystem
$ go run main.go wire_gen.go
```

If every it is ok, we must see:

```
Starting web server on port :8000
Starting gRPC server on port 50051
Starting GraphQL server on port 8080
```

### Call rest api POST /order

At the `api/create_order.http` file, send request `POST http://localhost:8000/order` just for test retrieving on the next calls.

### Call rest api GET /order

At the `api/create_order.http` file, send request `GET http://localhost:8000/order`.

### Call gRPC ListOrders

At g-expert/clean-architecture directory:

```
evans ./internal/infra/grpc/protofiles/order.proto
```

And at Evans prompt, call:

```
call ListOrders
```

### Call ListOrders GraphQL method

Open a browser and call the GraphQL playground at http://localhost:8080/. Then, at playground console, do the query:

```
query ListOrders {
  listOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```
