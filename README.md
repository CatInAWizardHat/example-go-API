# Getting Started
After cloning, run the following command to install dependancies
`go mod tidy`

Then, run the following command to build the docker image
`docker build . -t example-server`

Run with `docker run --rm --name go-example-server --network="host" --env-file .env example-server`

and access in the browser via `localhost:8080`
