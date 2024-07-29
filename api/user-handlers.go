package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Username string `json:"username" bson:"username"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")

	switch {
	case strings.HasPrefix(path, "users"):
		UserHandler(w, r)
	default:
		fmt.Fprintf(w, "Welcome to the main handler!")
	}
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	client := GetClient()
	if client == nil {
		http.Error(w, "Database connection not established", http.StatusInternalServerError)
		return
	}

	collection := client.Database("profileFolio").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var users []User
	if err = cursor.All(ctx, &users); err != nil {
		http.Error(w, "Error decoding users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
// 	id := strings.TrimPrefix(r.URL.Path, "/users/")

// 	// fmt.Printf("Requested user ID: %s\n", id) // Add this line for debugging

// 	collection := GetClient().Database("profileFolio").Collection("users")
// 	var user User
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	// If IDs are strings, use a string query instead of ObjectID
// 	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			http.Error(w, "User not found", http.StatusNotFound)
// 		} else {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 		}
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(user)
// }
