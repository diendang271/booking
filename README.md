# VN Bootcamp Services

Contains sample services 'bookings' and 'agents'

## Pre-requisites

In order to get started, make sure you have Golang, Docker and Protobuf installed:

1. Install Docker on MacOS (https://docs.docker.com/docker-for-mac/install/).
2. Install Protobuf through Homebrew (https://brew.sh/) by typing `brew install protobuf` in Terminal.
3. Install Golang via (https://golang.org/doc/install) or typing `brew install go` in Terminal. Make sure you have go modules enabled.

## Running the services

In order to run the services, simply open up the Terminal and call the script we have prepared for you.
```
./start.sh
```

This script takes care of building and deploying `bookings` service into a docker container along with other containers such as DynamoDB and `agents` service. If at any point you simply want to build, you can use the Makefile we've prepared.

```
make all
```

## Testing the bookings service

### createBooking

```
  curl http://localhost:8080/v1/createBooking -d ''
```

## Test bookings

Some sample data has already been created with the following BookingIDs. Use the dynamodb-admin tool below to view them.

```
test-booking-1
test-booking-2
test-booking-3
```

## dynamodb-admin

A [https://github.com/aaronshaf/dynamodb-admin](dynamodb-admin) server is included by default. Visit http://localhost:8001 to browser the local DynamoDB.
