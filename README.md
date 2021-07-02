# Golang Example API

![Test](https://github.com/Fantamstick/go-example-api/workflows/Test/badge.svg?branch=master)

## How to run

### Setup environment

Make sure that you have `direnv` installed to configure local environment variables. Please look at the [direnv github](https://github.com/direnv/direnv#install) for installation.
Copy the `.envrc.example` file and set the environment variables, then enable `direnv`.

```console
$ cp .envrc.sample .envrc
$ vi .envrc
$ direnv allow
```

And then, start the server and the MongoDB by runnig the command below:

```console
$ docker-compose up
```
