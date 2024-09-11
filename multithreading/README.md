# Multithreading

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
