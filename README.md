# Getting Started
After cloning, run the following command to install dependencies

```go mod tidy```

Then, run the following command to build the docker image

```docker build . -t example-server```

Run with 

```docker run --rm --name go-example-server --network="host" --env-file .env example-server```

Or build into a container with

```docker compose up --build```

## Endpoints:

### User Endpoints:
[GET] /users

[GET] /users/:id

### Message Endpoints:
[GET]    /messages

[GET]    /messages/:id

[POST]   /messages

[PATCH]  /messages/:id

[DELETE] /messages/:id


TODO:
- [ ] Finish defining user MemoryStore functions
- [ ] Finish defining user_handler endpoints
- [ ] Implement basic auth flow
- [ ] Add Unit tests for message_handler
- [ ] Finish integration tests for message_handler
- [ ] Add Unit tests for user MemoryStore
- [ ] Add Unit tests for user_handler endpoints
- [ ] Add integration tests for user_handler
