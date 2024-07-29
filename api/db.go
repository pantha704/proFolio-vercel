package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	once   sync.Once
)

func init() {
	initDB()
}

func initDB() {
	once.Do(func() {
		// Load local .env file
		// err := godotenv.Load()
		// if err != nil {
		// 	log.Fatalf("Error loading .env file")
		// }

		mongoURI := os.Getenv("MONGODB_URI")
		if mongoURI == "" {
			log.Println("MONGODB_URI environment variable is not set")
			return
		}

		clientOptions := options.Client().ApplyURI(mongoURI)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Printf("Error connecting to MongoDB: %v", err)
			return
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			log.Printf("Error pinging MongoDB: %v", err)
			return
		}

		log.Println("Connected to MongoDB successfully")
	})
}

func GetClient() *mongo.Client {
	if client == nil {
		// Attempt to reconnect
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Load local .env file
		// err := godotenv.Load()
		// if err != nil {
		// 	log.Fatalf("Error loading .env file")
		// }

		mongoURI := os.Getenv("MONGODB_URI")
		if mongoURI == "" {
			log.Println("MONGODB_URI environment variable is not set")
		}

		clientOptions := options.Client().ApplyURI(mongoURI)
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Println("Failed to reconnect to MongoDB:", err)
			return nil
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			log.Println("Failed to ping MongoDB after reconnection:", err)
			return nil
		}

		fmt.Println("Reconnected to MongoDB!")
	}
	return client
}

// Handler is a dummy handler to satisfy Vercel's requirements
func DBhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is a dummy handler for db.go")
}
