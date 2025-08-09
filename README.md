# Getting Started
After cloning, run the following command to install dependencies

```go mod tidy```

Then, run the following command to build the docker image

```docker build . -t example-server```

Run with 

```docker run --rm --name go-example-server --network="host" --env-file .env example-server```

Or build into a container with

```docker compose up --build```

and access in the browser via 

```localhost:8080/messages```

TODO:
[] Add Unit tests for handler
[] Finish integration tests for handler
