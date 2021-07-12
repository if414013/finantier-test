# Auth Service

Service that responsible to generate JWT and authenticate request to stock API given the token

## Framework

- Web : GoFiber
- Validation : Go-Ozzo
- Configuration : GoDotEnv

### Start the application

```bash
go run main.go
```

## Start the application

```bash
docker build -t auth-service .
docker run -d -p 9090:9090 auth-service
```