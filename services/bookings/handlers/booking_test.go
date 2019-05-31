// Copyright (c) 2012-2017 Grab Taxi Holdings PTE LTD (GRAB), All Rights Reserved. NOTICE: All information contained herein
// is, and remains the property of GRAB. The intellectual and technical concepts contained herein are confidential, proprietary
// and controlled by GRAB and may be covered by patents, patents in process, and are protected by trade secret or copyright law.
//
// You are strictly forbidden to copy, download, store (in any medium), transmit, disseminate, adapt or change this material
// in any way unless prior written permission is obtained from GRAB. Access to the source code contained herein is hereby
// forbidden to anyone except current GRAB employees or contractors with binding Confidentiality and Non-disclosure agreements
// explicitly covering such access.
//
// The copyright notice above does not evidence any actual or intended publication or disclosure of this source code,
// which includes information that is confidential and/or proprietary, and is a trade secret, of GRAB.
//
// ANY REPRODUCTION, MODIFICATION, DISTRIBUTION, PUBLIC PERFORMANCE, OR PUBLIC DISPLAY OF OR THROUGH USE OF THIS SOURCE
// CODE WITHOUT THE EXPRESS WRITTEN CONSENT OF GRAB IS STRICTLY PROHIBITED, AND IN VIOLATION OF APPLICABLE LAWS AND
// INTERNATIONAL TREATIES. THE RECEIPT OR POSSESSION OF THIS SOURCE CODE AND/OR RELATED INFORMATION DOES NOT CONVEY
// OR IMPLY ANY RIGHTS TO REPRODUCE, DISCLOSE OR DISTRIBUTE ITS CONTENTS, OR TO MANUFACTURE, USE, OR SELL ANYTHING
// THAT IT MAY DESCRIBE, IN WHOLE OR IN PART.

package handlers

import (
	"context"
	"testing"

	"github.com/micro/go-micro/client"
	"github.com/stretchr/testify/assert"
	cpb "gitlab.myteksi.net/michael.cartmell/bootcamp-skeleton/services/agents/pb"
	"gitlab.myteksi.net/michael.cartmell/bootcamp-skeleton/services/bookings/models"
	pb "gitlab.myteksi.net/michael.cartmell/bootcamp-skeleton/services/bookings/pb"
)

type mockStorer struct {
	updateCalled bool
}

type mockAgent struct{}

func (m *mockStorer) Save(interface{}) error {
	return nil
}

func (m *mockStorer) LoadBooking(bookingID string) (*models.Booking, error) {
	return &models.Booking{
		BookingID: "test_booking_id",
		State:     models.BookingState_InProgress,
	}, nil
}

func (m *mockStorer) UpdateBookingState(bookingID string, newState models.BookingState) error {
	m.updateCalled = true
	return nil
}

func (*mockAgent) NearbyAgent(ctx context.Context, in *cpb.NearbyAgentRequest, opts ...client.CallOption) (*cpb.NearbyAgentResponse, error) {
	return &cpb.NearbyAgentResponse{
		Agent: &cpb.AgentInfo{
			AgentId: "test_agent_id",
			Name:    "Test Agent",
		},
	}, nil
}

func testService() *BookingService {
	return &BookingService{
		store:  &mockStorer{},
		agents: &mockAgent{},
	}
}

func testLocation() *pb.Location {
	return &pb.Location{
		Latitude:  123,
		Longitude: 456,
	}
}

func TestCreateBooking(t *testing.T) {
	svc := testService()
	rsp := &pb.CreateBookingResponse{}
	err := svc.CreateBooking(context.Background(), &pb.CreateBookingRequest{
		Location: testLocation(),
	}, rsp)
	assert.Nil(t, err)
	assert.Equal(t, "test_agent_id", rsp.Booking.AgentId)
}

func TestGetBookingState(t *testing.T) {
	svc := testService()
	rsp := &pb.GetBookingStateResponse{}
	err := svc.GetBookingState(context.Background(), &pb.GetBookingStateRequest{
		BookingId: "test_booking_id",
	}, rsp)
	assert.Nil(t, err)
	assert.Equal(t, models.BookingState_InProgress, rsp.Booking.State)
	assert.Equal(t, "test_booking_id", rsp.Booking.BookingId)
}

func TestCompleteBooking(t *testing.T) {
	svc := testService()
	rsp := &pb.CompleteBookingResponse{}
	err := svc.CompleteBooking(context.Background(), &pb.CompleteBookingRequest{
		BookingId: "test_booking_id",
	}, rsp)
	assert.Nil(t, err)
	assert.True(t, svc.store.(*mockStorer).updateCalled)
}
