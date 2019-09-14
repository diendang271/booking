// Package handlers contains the API handlers.
package handlers

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-log/log"
	"github.com/micro/go-micro/client"
	agents_pb "gitlab.com/tech.learning.ext13/bootcamp-skeleton/services/agents/pb"
	"gitlab.com/tech.learning.ext13/bootcamp-skeleton/services/bookings/models"
	pb "gitlab.com/tech.learning.ext13/bootcamp-skeleton/services/bookings/pb"
	"os"
)

type AgentService interface {
	NearbyAgent(ctx context.Context, in *agents_pb.NearbyAgentRequest, opts ...client.CallOption) (*agents_pb.NearbyAgentResponse, error)
}

type BookingService struct {
	store  models.Storer
	agents AgentService
}

// New initializes the API handlers with dependencies
func New(store models.Storer, agentService agents_pb.AgentService) *BookingService {
	return &BookingService{
		store:  store,
		agents: agentService,
	}
}

func (s *BookingService) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest, rsp *pb.CreateBookingResponse) error {
	log.Log("Received BookingService.CreateBooking request")

	agentReq := agents_pb.NearbyAgentRequest{
		Location: &agents_pb.Location{
			Latitude:  req.Location.Latitude,
			Longitude: req.Location.Longitude,
		}}
	agentResp, err := s.agents.NearbyAgent(ctx, &agentReq, client.WithAddress("agent-service:9091"))
	if err != nil {
		fmt.Println(err)
	}

	ddb := newDDB()
	store := models.NewStore(ddb)
	booking := models.NewBooking()
	booking.Location = models.Location{Latitude: req.Location.Latitude, Longitude: req.Location.Longitude}
	booking.AgentID = agentResp.Agent.GetAgentId()
	booking.State = models.BookingState_InProgress

	err = store.Save(booking)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func newDDB() *dynamodb.DynamoDB {
	ddbURL := os.Getenv("DYNAMO_ENDPOINT")
	if ddbURL == "" {
		ddbURL = "http://localhost:8000"
	}
	log.Logf("using DDB endpoint %v", ddbURL)

	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("ap-southeast-1"),
		Endpoint: aws.String(ddbURL),
	})
	if err != nil {
		panic(err)
	}

	return dynamodb.New(sess)
}
