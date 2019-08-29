// Package handlers contains the API handlers.
package handlers

import (
	"context"

	"github.com/go-log/log"
	"github.com/micro/go-micro/client"
	agents_pb "gitlab.com/grab-learning-development/bootcamp-skeleton/services/agents/pb"
	"gitlab.com/grab-learning-development/bootcamp-skeleton/services/bookings/models"
	pb "gitlab.com/grab-learning-development/bootcamp-skeleton/services/bookings/pb"
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

// TODO: Implement handlers here
func (s *BookingService) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest, rsp *pb.CreateBookingResponse) error {
	log.Log("Received BookingService.CreateBooking request")
	return nil
}
