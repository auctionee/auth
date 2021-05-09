package db

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
)
var DBClient *firestore.Client

func CreateClient(ctx context.Context){
	projectID := "auctionee"

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	DBClient = client
}
