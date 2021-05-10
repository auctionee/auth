package db

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/auctionee/auth/pkg/data"
	"log"
)

var DBClient *firestore.Client

func CreateClient(ctx context.Context){

	client, err := firestore.NewClient(ctx, data.PROJECT_ID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	DBClient = client
}
func IfExist(ctx context.Context, creds *data.Credentials)error{
	login := creds.Login
	dsnap, err := DBClient.Collection("users").Doc("users").Get(ctx)
	if err != nil {
		return err
	}
	m := dsnap.Data()
	if _, ok:= m[login]; ok{
		return fmt.Errorf("login already exists")
	}
	return nil
}
func GetHashedPassword(ctx context.Context, creds *data.Credentials)string{
	login := creds.Login
	dsnap, err := DBClient.Collection("users").Doc("users").Get(ctx)
	if err != nil {
		return ""
	}
	m := dsnap.Data()
	if _, ok:= m[login]; !ok{
		return ""
	}
	return m[login].(string)
}