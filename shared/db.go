package shared

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
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
		err := godotenv.Load()
		if err != nil {
			log.Printf("Error loading .env file: %v", err)
		}

		mongoURI := os.Getenv("MONGODB_URI")
		if mongoURI == "" {
			log.Println("MONGODB_URI environment variable is not set")
			return
		}

		clientOptions := options.Client().ApplyURI(mongoURI)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// var err error
		client, err = mongo.Connect(ctx, clientOptions)
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
		initDB()
		if client == nil {
			log.Println("Error: Failed to initialize client.")
			return nil
		}
	}
	return client
}
