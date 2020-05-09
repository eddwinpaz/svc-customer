# Instructions

To be able to see what commands you can use on this customer service follow this command:

```
make help
```

Returns:

```
 Choose a command run in customer:

  dependency-prepare   download dep package manager in order to run make dependency
  dependency           Ensure all packages are installed
  coverage             run test coverage
  build                Clean build code as binary
  unittest             Run unit testing for code
  lint                 Run code linting and install golangci-lint in case it isn't present on folder
  docker-setup         Runs MySQL and PHPMyAdmin Docker containers
  setup-envs           Setup Enviroment variables
  run                  Run go server
  clean                Remove go build binary from folder
```

## Setup Dependencies

> Please make sure you have **Dependency management for Go** installed [download link](https://golang.github.io/dep/docs/installation.html)

Ensure dependency packages are installed

#### Install go dep package manager

```
  make setup-dependency
```

#### Install dependency packages

```
  make dependency
```

## Setup Database

1. Configure Database connection on **.env** file and add:

```
  MYSQL_ROOT_PASSWORD=root
  MYSQL_USER=root
  MYSQL_PASSWORD=root
  MYSQL_DATABASE=root
  MYSQL_HOST="127.0.0.1"
  MYSQL_PORT=3306
```

2. Use `customers.sql` and insert it on your database to migrate schema.

## Setup Enviroment Variables

Run command:

```
make setup-env
```

### Run Golang Server

Run command:

```
make run
```

## Test

To test the application code run command

```
make test
```

## Test / Coverage

To test the application code coverage run command

```
make test-coverage
```

# setup Docker Development containers

This will setup 2 containers containing a MySQL-Server Instance and a MySQL-PHPMyAdmin Instance; Remember to have docker installed.

```
make setup-docker
```

Now to check newly created docker containers run the following command:

```
docker ps
```

#### Redis Service

```
docker run --name redis-svc -p 6379:6379 -d redis
```

#### Redis Web UI

```
docker run --rm -ti -p 5001:5001 --link redis-svc:redis-root marian/rebrow
```

#### Distribute Git Hooks Between Teams.

Now, unfortunately the changes we made within the hooks/ directory under our project’s .git/ directory will not be tracked and therefore getting these changes out to various different members of your team becomes a bit of a challenge.

However, what you can do to get around this particular challenge is to create a directory called .githooks/ within your current project’s directory and store the pre-commit git hook within that directory. You’ll be able to commit and track this just as you would any other files within your project and in order to turn on these enable these hooks on other development machines you simply need to run this command:

```
git config core.hooksPath .githooks
```