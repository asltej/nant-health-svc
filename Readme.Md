## NANT Health Service is a Coding challenge for Nant Health

## Technical Stack

Language: Golang

API: RESTFul

Database: Inmemory Slice.

## How to Run

```
go run cmd/main.go
```

## How to Test

There are two endpoints in this services.

1) Post Metrics.

```
curl --location --request POST 'localhost:8080/v1/metrics' \
--header 'Content-Type: application/json' \
--data-raw '{
    "machineId": 12345,
    "stats": {
        "cpuTemp": 90,
        "fanSpeed": 400,
        "HDDSpace": 800
    },
    "lastLoggedIn": "admin/Paul",
    "sysTime": "2022-04-23T18:25:43.511Z"
}'
```

1. Get Metrics.

```
curl --location --request GET 'localhost:8080/v1/metrics'
```

## How to check the code coverage

Run `go test -cover ./...` command in terminal, which will result in showing the coverage
