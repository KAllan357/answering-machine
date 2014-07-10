# Answering Machine

A dead simple Go application that starts a server on 8080 and has does two things:

* Accepts a POST to the /messages endpoint.
* Accepts a GET to the /messages endpoint.

This repository will serve as a simple "training ground" for Docker. Answering Machine will have a component application,
that will POST messages to the /messages endpoint. This repository will have an intentionally "weak" Dockerfile, which has just enough
information for a noob to get started with running their own Answering Machine.

You should leave a message!

# Installation

You'll need an installation of Go, and once Go is on your path:

```
go run main.go
```

This will start the server. You should get a 200 response if you:

```
curl -v localhost:8080/messages
```

# The Exercise

1. Start a container with this Dockerfile.
2. I want to curl the Answering Machine, how do I expose it to my localhost?
3. This is my answering machine, not Docker's! How can I change the ANSWERING_MACHINE_OWNER environment variable so I can make this my own?
4. 