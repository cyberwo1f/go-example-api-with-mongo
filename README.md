# Golang Example API with MongoDB

![Test](https://github.com/cyberwo1f/go-example-api/workflows/Test/badge.svg?branch=master)

This repository is a template showing the basic architecture and file system for developing APIs using MongoDB in the Go language.

The architecture references the following design patterns to maintain loose coupling, domain management efficiency, and tuning cost performance.

- DDD
- clean architecture

This is not a solution for every need, but only a basic architecture.
It should be rearranged as needed depending on your needs, but note that the relationships between entities, repositories, persistence, and other injections, as well as the scope of responsibility and processes for each package, should be maintained.

## How to run

### Setup environment

Make sure that you have `direnv` installed to configure local environment variables. Please look at the [direnv GitHub](https://github.com/direnv/direnv#install) for installation.
Copy the `.envrc.example` file and set the environment variables, then enable `direnv`.

```console
$ cp .envrc.sample .envrc
$ vi .envrc
$ direnv allow
```

And then, start the server and the MongoDB by running the command below:

```console
$ docker-compose up
```

## API Interfaces

| Func        | Method | Path                   | Description                            |
|-------------|--------|------------------------|----------------------------------------|
| GetUsers    | GET    | /user/list             | get list by all users                  |
| GetMessages | GET    | /message/list/:user_id | get all messages for specified user id |
