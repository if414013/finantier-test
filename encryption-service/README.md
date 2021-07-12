# Encryption Service

Service that responsible to encrypt Stock Information Data

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
docker build -t encryption-service .
docker run -d -p 7070:7070 encryption-service
```