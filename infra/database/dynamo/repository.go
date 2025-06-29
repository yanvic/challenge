package dynamo

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
)

// var Client *dynamodb.Client

var tableGps = "GPSTable"
var tableGyro = "GyroscopeTable"
var tablePhoto = "PhotoTable"

func SaveGps(latitude float64, longitude float64, deviceID string, timestamp string, payload interface{}) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	println("Payload JSON:", string(jsonData))

	lat := latitude
	lon := longitude

	item := GpsItem{
		UUID:      uuid.NewString(),
		Latitude:  &lat,
		Longitude: &lon,
		DeviceID:  deviceID,
		Timestamp: timestamp,
	}
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return err
	}

	_, err = Client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: &tableGps,
		Item:      av,
	})
	return err
}

func SaveGyro(x float64, y float64, z float64, deviceID string, timestamp string, payload interface{}) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	println("Payload JSON:", string(jsonData))

	pitch := x
	yaw := y
	roll := z

	item := GyroscopeItem{
		UUID:      uuid.NewString(),
		X:         &pitch,
		Y:         &yaw,
		Z:         &roll,
		DeviceID:  deviceID,
		Timestamp: timestamp,
	}
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return err
	}

	_, err = Client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: &tableGyro,
		Item:      av,
	})
	return err
}

func SavePhoto(ImageBase64 string, deviceID string, timestamp string, payload interface{}) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	println("Payload JSON:", string(jsonData))

	item := PhotoItem{
		UUID:        uuid.NewString(),
		ImageBase64: ImageBase64,
		DeviceID:    deviceID,
		Timestamp:   timestamp,
	}
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return err
	}

	_, err = Client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: &tablePhoto,
		Item:      av,
	})
	return err
}
