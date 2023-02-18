# GOGOLF Boilerplate Template

My go backend template. This is a go boiplerplate to quickly started the new project with golang and standard dependencies for local development pipeline.
![Tux, the Linux mascot](/docs/GOGolf.jpg)

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

**Table of Contents**

- [Getting Started](#getting-started)
- [Prerequsition](#prerequsition)
- [👩‍⚕️ Pre-Commit](#%E2%80%8D-pre-commit)
- [Commit Lint](#commit-lint)
- [Miscellaneous](#miscellaneous)
- [Set $GOPATH](#set-gopath)
- [Hexagonal Architecture](#hexagonal-architecture)
- [Architecture Components](#architecture-components)
- [Core](#core)
- [Domain](#domain)
- [Service](#service)
- [Repository](#repository)
- [Dependency Injection](#dependency-injection)
  - [Port and Adaptor Design Patterns](#port-and-adaptor-design-patterns)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

**Design Structure**
[สรุปเรื่อง GopherCon 2018: Kat Zien — How Do You Structure Your Go Apps](https://goangle.medium.com/%E0%B8%AA%E0%B8%A3%E0%B8%B8%E0%B8%9B%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87-gophercon-2018-kat-zien-how-do-you-structure-your-go-apps-a96faca9e8f1)

**Features**

1. Hot-Reloader with Air.
2. Pre-commit.

# Getting Started

## Prerequsition

Please run `$ make init` to initialize the basic needed command and dependencies for your local development. We built this command to install the stuff to make the industry stnadard good grade to local development.

Here is a standard development pipeline installed.

1. go
2. node
3. pre-commit

## 👩‍⚕️ Pre-Commit

You can reached pre-commit hooks and rules under `.pre-commit-config.yaml` configuration file.

## Commit Lint

The commit lint is make more sense of good engineering should be. Essepcially, the `commitlint`, in this template I've follow the convention here [https://www.conventionalcommits.org/en/v1.0.0/](https://www.conventionalcommits.org/en/v1.0.0/)

## Miscellaneous

### Set $GOPATH

The $GOPATH environment variable lists places for Go to look for Go Workspaces.

[https://go.dev/doc/gopath_code#GOPATH](https://go.dev/doc/gopath_code#GOPATH)

Add this to your profile. If you use _zsh_ please also update `.zshrc`

```sh
export GOPATH="$HOME/go"
PATH="$GOPATH/bin:$PATH"
```

then `source .zshrc` for reload the the environment. You can check you GOPATH by `go env GOPATH`

# Hexagonal Architecture

This golang project adopt the hexagonal architecture to grouping the code by domain and context of business problems.

## Architecture Components

### Core

Everything surrounded the `core` this layer contains the concrete core business logics.

The core could be viewed as a “box” (represented as a hexagon) capable of resolve all the business logic independently.

### Domain

The domain entity it contains the go struct definition of each entity that is part of the domain problem and can be used across the application.

> **Seperate of concern**
> For example, this layer should not be known how the data keeps its, or how the caching store in the redis.

**Port** is the interfaces are the ports that external actors will plug their adapters into to drive the application.

**Adapter** is the implementation of the ports

### Service

Service is through which the client will interact with actual business logic/domain objects. The client can be rest handlers, test agents, etc.

### Repository

The repository interface allows us to connect to multiple data sources.

### Dependency Injection

#### Port and Adaptor Design Patterns

Outside concerns, such as repositories and user interfaces, use the adapters to plug into the business domain services.

**References**

- https://www.linkedin.com/pulse/part-2-go-project-layout-hexagonal-architecture-kaushal-prajapati/
- https://www.linkedin.com/pulse/hexagonal-software-architecture-implementation-using-golang-ramaboli/
- https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3

<!-- ------------------------------- -->

## GOLANG TODOS APPLICATION

### Technical stuff

- Architecture: Clean architecture
- Framework: Echo
- ORM: Gorm
- DB: Postgres
- Deployment: Docker

### Overview

- Support JWT
- Limit todos per user in a day
- Unit tests & integration test

### What next ?

- Add role/permission based validation
- Add missing tests
- Implement your new features

### How to run the code locally

Clone the project then:

##### Update .env file

```txt
PORT=8080
JWT_SECRET=B5bJHoI8aVLjAAeV
SIGNING_KEY=ABC
HASH_SALT=SJSHDFDS
TOKEN_TTL=86400
CONNECTION_URL=host=localhost user=postgres password=password1 dbname=todos port=5432
```

```bash
go run cmd/api/main.go
```

##### or by Docker

Update .env file (change host to postgresql)

```txt
#...
CONNECTION_URL=host=postgresql user=postgres password=password1 dbname=todos port=5432
```

RUN COMMAND:

```bash
docker-compose up -d
```

### Then open postman or bash use curl:

#### Register:

```bash
curl --location --request POST 'http://localhost:8080/api/v1/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "kiendinh",
    "password": "abc123",
    "limit": 2
}'
```

#### Login:

```bash
curl --location --request POST 'http://localhost:8080/api/v1/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "kiendinh",
    "password": "abc123"
}'
```

Take the token from login then.

#### Add todo:

```bash
curl --location --request POST 'http://localhost:8080/api/v1/todos/' \
--header 'Authorization: Bearer YOUR_TOKEN_HERE' \
--header 'Content-Type: application/json' \
--data-raw '{
    "content": "TEST TO DO 2x"
}'
```

#### Get all todos:

Public for all users

```bash
curl --location --request GET 'http://localhost:8080/api/v1/todos/'
```

#### Get user todos:

Get any user's todos with their id

```bash
curl --location --request GET 'http://localhost:8080/api/v1/todos/A_USER_ID' \
--header 'Authorization: Bearer YOUR_TOKEN_HERE'
```

### TESTS

I just implemented some necessary test samples, we can add tests of handlers, usecases and another endpoints ...

#### UNIT TEST

Command:

```bash
go test ./...
```

#### INTEGRATION TEST

Command:

```bash
go test -v ./integration-tests
```

### Conclusion

First time I created a new microservice with Go from scratch, It gave me a challenge but I did it, tried to remove the old mindset in another architecture then go to clean architecture. I love it :heart: :heart: .
