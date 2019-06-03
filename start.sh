#!/bin/bash
docker-compose up --build -d
DDB="docker-compose run -e AWS_ACCESS_KEY_ID=A -e AWS_SECRET_ACCESS_KEY=B awscli aws dynamodb"
$DDB create-table --table-name bootcamp_bookings \
  --region ap-southeast-1 \
  --key-schema AttributeName=BookingID,KeyType=HASH \
  --attribute-definitions AttributeName=BookingID,AttributeType=S \
  --provisioned-throughput ReadCapacityUnits=100,WriteCapacityUnits=100 \
  --endpoint-url http://storage:8000
docker-compose ps

# Create test bookings
DDB_PUT="$DDB put-item --table-name bootcamp_bookings --region ap-southeast-1 --endpoint-url http://storage:8000"
$DDB_PUT --item '{"BookingID": {"S": "test-booking-1"}, "Location": {"M": {"Latitude": {"N":"1.3521"}, "Longitude":{"N":"103.8198"}}}, "AgentID":{"S": "test-agent-1"}, "UserID":{"S": "user-id-1"}, "State":{"S": "NEW"}}'
$DDB_PUT --item '{"BookingID": {"S": "test-booking-2"}, "Location": {"M": {"Latitude": {"N":"1.3521"}, "Longitude":{"N":"103.8198"}}}, "AgentID":{"S": "test-agent-2"}, "UserID":{"S": "user-id-1"}, "State":{"S": "IN_PROGRESS"}}'
$DDB_PUT --item '{"BookingID": {"S": "test-booking-3"}, "Location": {"M": {"Latitude": {"N":"1.3521"}, "Longitude":{"N":"103.8198"}}}, "AgentID":{"S": "test-agent-3"}, "UserID":{"S": "user-id-1"}, "State":{"S": "COMPLETED"}}'
echo 'Started the required containers, to stop run "./stop.sh"'
