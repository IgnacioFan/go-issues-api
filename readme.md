# go-issues-api
Go Issues API is an API server for raising issues, voting and viewing voting issues. This repo is also a demonstration of a scalable system with Golang.

## Introduction
The structure of the entire repo buys in the concept of the clean architecture and refers to a couple of great materials such as [go-clean-arch](https://github.com/bxcodec/go-clean-arch), [Trying Clean Architecture on Golang](https://medium.easyread.co/golang-clean-archithecture-efd6d7c43047) and [Applying The Clean Architecture to Go applications](https://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/). This repo adopted the four following layers.

- **Model layer** as entities which contains business rules.
- **Usecase layer** as usecase which proceeds the business logic. Data will be fetched at the repository layer, the usecase layer is responsible for processing data based on the business requirements.
- **Repository layer** as interface adapters that interacts with database. This layer supports the operation of the database. For example, Querying, and inserting data will be processed here.
- **Delivery layer** as interface adapters that works as the presenter to decide how data will be structured and presented to client. In terms of web API, the presenter can support various types such as REST API, gRPC or CLI etc.
## System overview
TBD

## How to run and test?

### Run the application
After cloning the repo at your workspace, run `make app.start` to run services with docker-compose.
No need to worry about data migrations, because the Application will conduct DB migration during the initialization.

```
git clone https://github.com/IgnacioFan/go-issues-api.git

make app.start
```

To simply test the connection, you can use curl or test full API endpoints via Postman.
For checking the available API endpoints, please take a look [Available API endpoints](#available-api-endpoints).

```
curl -X GET http://localhost:8080/api/v1/ping
```

You should see the response below.
```
"success"
```

### Run unit test
Perform unit test:
```
make test.unit
```

## Available API endpoints
TBD

## Tools introduction
1. https://github.com/golang-migrate/migrate
  - to generate migration files and conduct migrate.up and migrate.down
2. https://github.com/vektra/mockery
  - to generate mocks file for http and repository testing.

## What's Next?
