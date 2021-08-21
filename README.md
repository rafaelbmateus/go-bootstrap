# go-bootstrap

[Golang](https://golang.org)
boostrap project using
[docker](https://docker.com),
[docker-compose](https://docs.docker.com/compose) and
[makefile](https://www.gnu.org/software/make/manual/make.html).

```
# Build and start development environment in background.
make up

# Follows development logs [service="svc1 svc2..."].
make logs

# Start a shell session within the container.
make shell

# Runs static analysis code.
make lint

# Runs the tests.
make test

# Stop development environment.
make clean
```
