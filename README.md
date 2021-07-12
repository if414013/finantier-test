# Simple Stock Info Retrieval Microservices
Notes : 
    - Really sorry if I don't implement TDD, the timeline to tight, only have 1 days :(
    - I implement encrypt in fields level
    - Need to generate token first before calling getStockInfoAPI
    - Token will be expired after 10 minutes
## High Level Diagram

![diagram](image/diagram.png)

- [Stock Service](stock-service/README.md)
- [Auth Service](auth-service/README.md)
- [Encryption Service](encryption-service/README.md)

Postman Collection can be found [HERE](Postman_Collection.json)

## How to Run

```bash
docker compose up
```

## Attachments

#### Server Running

![diagram1](image/running.png)

#### Get Token API

![diagram2](image/get_token.png)

#### Get Stock Info Invalid Token

![diagram3](image/get_stock_info_token_expired.png)

#### Get Stock Info Invalid Stock Symbol

![diagram4](image/get_stock_info_invalid_stock_name.png)

#### Get Stock Info Success

![diagram5](image/get_stock_info_success.png)

![diagram6](image/get_stock_info_console.png)