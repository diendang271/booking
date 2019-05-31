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

package models

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-log/log"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	testInit()
	os.Exit(m.Run())
}

func testInit() {
	os.Setenv("AWS_ACCESS_KEY_ID", "A")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "B")
	os.Setenv("AWS_REGION", "ap-southeast-1")
}

func testDDB() *dynamodb.DynamoDB {
	ddbURL := os.Getenv("DYNAMO_ENDPOINT")
	if ddbURL == "" {
		return nil
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

func TestBooking(t *testing.T) {
	ddb := testDDB()
	if ddb == nil {
		t.Log("test skipped -- set DYNAMO_ENDPOINT")
		t.Skip()
		return
	}
	store := NewStore(ddb)
	booking := NewBooking()
	booking.AgentID = "test"

	err := store.Save(booking)
	assert.Nil(t, err)

	err = store.UpdateBookingState(booking.BookingID, BookingState_Completed)
	assert.Nil(t, err)

	res, err := store.LoadBooking(booking.BookingID)
	assert.Nil(t, err)
	assert.Equal(t, "test", res.AgentID)
}
