# Bookings Service

This is the skeleton Bookings service

## Getting Started

- [Usage](#usage)

## Structure

```
  gateway/  - The grpc-gateway code (HTTP proxy)
  handlers/ - The API controllers
  models/   - The database store and entitieS
  pb/       - The proto definition
```

## Usage

A Makefile is included for convenience

Build the proto

```
make proto
```

Build the binary

```
make build
```

Run the service

```
./bookings-srv --server_address=localhost:9090
```

Test the service (via gateway on port :8080)

```
curl http://localhost:8080/v1/createBooking -d ''
```

Build a docker image

```
make docker
```
