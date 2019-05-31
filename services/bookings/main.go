// Package main is the entrypoint for the bookings service.
package main

import (
	"context"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	mgrpc "github.com/micro/go-grpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	agentspb "gitlab.myteksi.net/michael.cartmell/bootcamp-skeleton/services/agents/pb"
	gwpb "gitlab.myteksi.net/michael.cartmell/bootcamp-skeleton/services/bookings/gateway/pb"
	"gitlab.myteksi.net/michael.cartmell/bootcamp-skeleton/services/bookings/handlers"
	"gitlab.myteksi.net/michael.cartmell/bootcamp-skeleton/services/bookings/models"
	pb "gitlab.myteksi.net/michael.cartmell/bootcamp-skeleton/services/bookings/pb"
	"google.golang.org/grpc"
)

func main() {
	// Construct service
	service := micro.NewService(
		micro.Name("bootcamp.bookings"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()
	service.Server().Init(
		server.Address("0.0.0.0:9090"),
	)

	// Register handlers
	store := models.NewStore(newDDB())
	hdlr := handlers.New(store, newCClient())
	pb.RegisterBookingServiceHandler(service.Server(), hdlr)

	// Run grpc-gateway
	go runGateway()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

// Dependencies

// newCClient creates a new client for the agent service. Note it uses service discovery, so doesn't need the host.
func newCClient() agentspb.AgentService {
	service := mgrpc.NewService()
	return agentspb.NewAgentService("bootcamp.srv.agents", service.Client())
}

// newDDB creates a new connection to the local DynamoDB. Uses $DYNAMO_ENDPOINT by default, otherwise
// it uses http://localhost:8000.
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

// runGateway starts a grpc-gateway on port :8080
func runGateway() {
	endpoint := "localhost:9090"
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := gwpb.RegisterBookingServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", mux)
}
