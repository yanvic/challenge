package main

import (
	"context"
	"log"

	"github.com/yanvic/challenge/infra/database/dynamo"
)

func main() {
	ctx := context.TODO()

	client, err := dynamo.InitDynamoClientTest(ctx)
	if err != nil {
		log.Fatalf("failed to initialize dynamodb client: %v", err)
	}

	dynamo.Client = client

	if err := dynamo.CreateTable(ctx, client, "GyroscopeTable"); err != nil {
		log.Fatalf("error creating GyroscopeTable: %v", err)
	}

	if err := dynamo.CreateTable(ctx, client, "GPSTable"); err != nil {
		log.Fatalf("error creating GPSTable: %v", err)
	}

	if err := dynamo.CreateTable(ctx, client, "PhotoTable"); err != nil {
		log.Fatalf("error creating PhotoTable: %v", err)
	}

	if err := dynamo.CreateTable(ctx, client, "PhotoAnalysisTable"); err != nil {
		log.Fatalf("error creating PhotoAnalysisTable: %v", err)
	}
}
