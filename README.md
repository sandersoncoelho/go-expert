# Go-expert

This repository contains implementations for challenges proposed at the go-expert postgraduate. The challenges are named as:

- [Client-Server-API](#client-server-api)
- [Multithreading](#multithreading)
- [Clean-Architecture](https://github.com/sandersoncoelho/go-expert/tree/master/clean-architecture)

## Versions

- Go: 1.22.1
- SQLite3: 3.42.0

## Client-Server-API

This is a simple example for a client and server application written with Go language. The client sends a request to server that requests an external API, so the server processes the data and return to client.

In this example it was used some concepts such as: http, context, json, panic recover, database persistence and writing on files.

### Database

The database cotacao.db was commited empty and It was created at the client-server-api directory as follow:

```
$ sqlite3 cotacao.db
$ sqlite> CREATE TABLE COTACAO(ID INTEGER PRIMARY KEY AUTOINCREMENT, BID REAL);
```

### Running the applications

```
~/go-expert cd client-server-api
~/client-server-api go run server/server.go
~/client-server-api go run client/client.go
```

### Comments

At the server application we can see the message "Erro na requesição para economia.awesomeapi.com.br", because the context has 200ms to complete the request to economia.awesomeapi.com.br, and this is not enough for that. On the client side, we get "Get "http://localhost:8080/cotacao": context deadline exceeded", explaining to client what happened at the server side.
Thereby, one option is to increase the timeout at the contexts, such as 2000ms at server and 3000ms at client. These values can be adjust arbitrarily adapting to the environment performance.

## Multithreading

In this challenge we have an endpoint that called two external service, cepbrasil and viacep. The first service response to return, will be showed to the client, and the other service response will be discarded. We use go routines to called the both services simultaneously.

### Running the applications

Run the server in a terminal

```
~/go-expert cd multithreading
~/multithreading go run cmd/server/main.go
```

Call the endpoint in other terminal passing the cep at the request http://localhost:8000/ceps/{cep_query}

```
curl  http://localhost:8000/ceps/01153000
```

### Comments

The response to the client must return which service has returned first as well its data, such as:

```
{
   "serviceName":"brasilapi",
   "data":{
      "cep":"01153000",
      "city":"São Paulo",
      "neighborhood":"Barra Funda",
      "street":"Rua Vitorino Carmilo",
      "service":"open-cep"
   }
}
```
