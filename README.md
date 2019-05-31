# VN Bootcamp Services

Contains sample services 'bookings' and 'agents'. See https://wiki.grab.com/pages/viewpage.action?pageId=242545747

## Running the services

```
  make
  ./start.sh
```

## Testing the bookings service

### createBooking

```
  curl http://localhost:8080/v1/createBooking -d '{"location":{"latitude":1.3521, "longitude": 103.8198}}'
```

### getBookingState

```
  curl http://localhost:8080/v1/getBookingState?booking_id=388b53c959366a536f13745bcbc13e77
```

### completeBooking

```
  curl http://localhost:8080/v1/completeBooking -d '{"booking_id" "e106c895813eb88062a023f5b6ff7160"}'
```

## dynamodb-admin

A [https://github.com/aaronshaf/dynamodb-admin](dynamodb-admin) server is included by default. Visit http://localhost:8001 to browser the local DynamoDB.
