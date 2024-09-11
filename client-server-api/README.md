# Client-Server-API

This is a simple example for a client and server application written with Go language. The client sends a request to server that requests an external API, so the server processes the data and return to client.

In this example it was used some concepts such as: http, context, json, panic recover, database persistence and writing on files.

### Database

It was used SQLite3, 3.42.0 version. The database cotacao.db was commited empty and It was created at the client-server-api directory as follow:

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
