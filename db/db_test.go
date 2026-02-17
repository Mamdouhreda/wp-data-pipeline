package db

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func TestConnectDB(t *testing.T) {
	client, err := ConnectToDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	defer client.Disconnect(context.TODO())

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}

	_ = client.Database(DatabaseName)
	t.Log("Connected to database successfully")
}
