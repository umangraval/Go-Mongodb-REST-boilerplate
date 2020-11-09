# Go-Mongodb-REST-boilerplate


![Build](https://github.com/umangraval/Go-Mongodb-REST-boilerplate/workflows/Go/badge.svg)


This repo can be used as a starting point for backend development with Golang. It comes bundled with Docker. The development environment uses `docker-compose` to start dependent services like mongo.

A few things to note in the project:
* **[Github Actions Workflows](https://github.com/umangraval/Go-Mongodb-REST-boilerplate/tree/main/.github/workflows)** - Pre-configured Github Actions to run automated builds and publish image to Github Packages
* **[Dockerfile](https://github.com/umangraval/Go-Mongodb-REST-boilerplate/blob/main/Dockerfile)** - Dockerfile to generate docker builds.
* **[docker-compose](https://github.com/umangraval/Go-Mongodb-REST-boilerplate/blob/main/docker-compose.yml)** - Docker compose script to start service in production mode.
* **[Containerized Mongo for development](#development)** - Starts a local mongo container with data persistence across runs.
* **[Mongo Driver](https://go.mongodb.org/mongo-driver)** - MongoDB supported driver for Go.
* **[Gorilla Mux](https://go.mongodb.org/mongo-driver)** - HTTP request multiplexer.
* **[jwt-go](https://github.com/dgrijalva/jwt-go)** - Implementation of JWT Tokens.
* **[Validator](https://gopkg.in/go-playground/validator.v9)** - Package validator implements value validations for structs.
<!-- * **[OpenAPI 3.0 Spec](https://github.com/sidhantpanda/docker-express-typescript-boilerplate/blob/master/openapi.json)** - A starter template to get started with API documentation using OpenAPI 3.0. This API spec is also available when running the development server at `http://localhost:3000/dev/api-docs` -->
* **[.env file for configuration](#environment)** - Change server config like app port, mongo url etc
* **[File Uploads](https://golang.org/pkg/io/)** - io package provides interfaces to I/O primitives.
* **[httptest](#testing)** - Utilities for HTTP testing.

## Installation

#### 1. Clone this repo

```
$ git clone git@github.com:umangraval/Go-Mongodb-REST-boilerplate.git your-app-name
$ cd your-app-name
```

#### 2. Install dependencies

```
$ go mod vendor
```

## Development

### Start dev server
Starting the dev server also starts MongoDB as a service in a docker container using the compose script at `docker-compose.yml`.

```
$ go run main.go routes.go
```
Running the above commands results in 
* 🌏 **API Server** running at `http://localhost:8080`
<!-- * ⚙️**Swagger UI** at `http://localhost:3000/dev/api-docs` -->
* ⛁ **MongoDB** running at `mongodb://localhost:27017/db`

## Packaging and Deployment
#### 1. Build and run without Docker

```
$ go build 
```
#### 2. Run Tests

```
$ cd tests
$ go test
```
#### 3. Run with docker

```
$ docker build -t api-server .
$ docker run -t -i -p 8080:8080 api-server
```

#### 4. Run with docker-compose

```
$ docker-compose up
```


---

## Environment
To edit environment variables, create a file with name `.env` and copy the contents from `.env.default` to start with.

| Var Name  | Type  | Default | Description  |
|---|---|---|---|
| JWT_SECRET  | string  | `secret` |JWT secret to verify  |
|  PORT | number  | `8080` | Port to run the API server on |
|  MONGO_URL | string  | `mongodb://localhost:27017/db` | URL for MongoDB |

<!-- ## Logging
The application uses [winston](https://github.com/winstonjs/winston) as the default logger. The configuration file is at `src/logger.ts`.
* All logs are saved in `./logs` directory and at `/logs` in the docker container.
* The `docker-compose` file has a volume attached to container to expose host directory to the container for writing logs.
* Console messages are prettified
* Each line in error log file is a stringified JSON. -->


### Directory Structure

```
+-- controllers
|   +-- personController.go
+-- db
|   +-- db.go
+-- handlers
|   +-- config.go
|   +-- logs.go
|   +-- response.go
|   +-- verifyJWT.go
+-- models
|   +-- models.go
+-- validators
|   +-- validators.go
+-- tests
|   +-- api_test.go
+-- routes
|   +-- routes.go
+-- uploaded
+-- vendor
+-- nginx
|   +-- dev.conf.d
|   |   +-- nginx.conf
+-- .env
+-- .env.default
+-- .gitignore
+-- docker-compose.yml
+-- Dockerfile
+-- go.mod
+-- go.sum
+-- main.go
+-- README.md
```