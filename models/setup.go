package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// const connectionString = mongo_db_3_connection_string
// const connectionString = mongo_db_6_connection_string
// const connectionString = mongo_db_7_connection_string

const db = "test"
const collection = "users"

var mongoClient *mongo.Client;

func ConnectDatabase(){
	fmt.Println("Connecting to the database...");

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions);
	if err != nil {
		panic(err)
	}

	var result map[string]interface{}
	err = client.Database("admin").RunCommand(context.TODO(), map[string]interface{}{"buildInfo": 1}).Decode(&result)
	if err != nil {
		fmt.Println("Failed to run buildInfo command:", err)
	}

	// Print the version
	if version, ok := result["version"]; ok {
		fmt.Println("MongoDB Server Version:", version, "\n\n");
	} else {
		fmt.Println("Could not find version in response")
	}

	mongoClient = client
}