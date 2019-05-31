# VN Bootcamp Services

Contains sample services 'bookings' and 'agents'

## Running the services

```
  make
  ./start.sh
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
