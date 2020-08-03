// shippy-service-consignment/datastore.go
package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateClient -
func CreateClient(ctx context.Context, uri string, retry int32) (*mongo.Client, error) {
	log.Print("Attempting to connect to MongoDB.")
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err := conn.Ping(ctx, nil); err != nil {
		log.Printf("Error connecting to Mongo DB: %s", err)
		if retry >= 3 {
			log.Printf("Retry count of %s exceeded. Exiting", retry)
			return nil, err
		}
		retry = retry + 1
		time.Sleep(time.Second * 2)
		return CreateClient(ctx, uri, retry)
	}

	return conn, err
}
