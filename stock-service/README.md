# Stock Service

Service that responsible to serve request from client and call Alphavantage API

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
docker build -t stock-service .
docker run -d -p 8080:8080 stock-service
```