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
//
// Package models contains the database storage for bookings.
package models

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	tableName = "bootcamp_bookings"
)

type ModelStore struct {
	db *dynamodb.DynamoDB
}

type Storer interface {
	Save(interface{}) error
	LoadBooking(bookingID string) (*Booking, error)
	UpdateBookingState(bookingID string, newState BookingState) error
}

func NewStore(db *dynamodb.DynamoDB) *ModelStore {
	return &ModelStore{
		db: db,
	}
}

// LoadBooking retrieves a booking from dynamodb
func (m *ModelStore) LoadBooking(bookingID string) (*Booking, error) {
	data := &Booking{}
	out, err := m.db.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"BookingID": {
				S: aws.String(bookingID),
			},
		},
		TableName: aws.String(tableName),
	})
	if err != nil {
		return nil, err
	}
	if out.Item == nil {
		return nil, errors.New("booking not found")
	}
	if err = dynamodbattribute.UnmarshalMap(out.Item, data); err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateBookingState updates an existing booking with a new state
func (m *ModelStore) UpdateBookingState(bookingID string, newState BookingState) error {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":s": {
				S: aws.String(string(newState)),
			},
		},
		ExpressionAttributeNames: map[string]*string{
			"#S": aws.String("State"),
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"BookingID": {
				S: aws.String(bookingID),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String(`set #S = :s`),
	}
	_, err := m.db.UpdateItem(input)
	return err
}

// Save saves arbitrary data to the table
func (m *ModelStore) Save(data interface{}) error {
	av, err := dynamodbattribute.MarshalMap(data)
	if err != nil {
		return err
	}
	putInput := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}
	_, err = m.db.PutItem(putInput)
	return err
}
