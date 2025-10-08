package models

import (
	"fmt"
	"encoding/json"
	"context"
	"reflect"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertOne() string {
	fmt.Print("---------------------------------\nInsert One\n---------------------------------\n")
	collection := mongoClient.Database(db).Collection(collection)
	inserted, err := collection.InsertOne(context.TODO(), newUser)
	jsonData, err := json.MarshalIndent(inserted, "", " ")
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Result: ", string(jsonData))
	fmt.Println()
	fmt.Println("Return type:", reflect.TypeOf(inserted)); // 6: *mongo.InsertOneResult
	fmt.Println("Return type of InsertedID", reflect.TypeOf(inserted.InsertedID)); //6: primitive.ObjectID
	fmt.Println();
	return inserted.InsertedID.(primitive.ObjectID).Hex()
}


func InsertMany() error {
	fmt.Print("---------------------------------\nInsert Many\n---------------------------------\n")
	newUsers := make([]interface{}, len(users));
	for i, user := range users {
		newUsers[i] = user;
	}

	collection := mongoClient.Database(db).Collection(collection);
	result, err := collection.InsertMany(context.TODO(), newUsers);
	if err != nil {
		fmt.Println("Error:", err)
	}
	jsonData, err := json.MarshalIndent(result, ""," ");
	if err != nil {
		fmt.Println("Marshal Error:", err)
	}

	fmt.Println("Result:", string(jsonData));
	fmt.Println()
	fmt.Println("Return type:", reflect.TypeOf(result))
	fmt.Println();
	return err
}

func InsertManyUnordered() {
	fmt.Print("---------------------------------\nInsert Many Unordered - Bulk Run\n---------------------------------\n")
	// put a duplicate id here to see that
	// unordered would skip the error and will add the remaining inserts
	// without crashing or leaving in between
	// the one with duplicate id wouldn't appear in the result/DB.
	idd, _ := primitive.ObjectIDFromHex("68e4dc2c176e1709eaaade0d")
	userss := []bson.M{
		{"_id": idd, "name": "test10", "age": 44, "balance": 3433454},
		{"name": "test20", "age": 80, "balance": 3343},
		{"name": "test30", "age": 57, "balance": 5000},

	}

	newUsers := make([]interface{}, len(userss));
	for i, user := range userss {
		newUsers[i] = user;
	}

	opts := options.InsertMany().SetOrdered(false);

	collection := mongoClient.Database(db).Collection(collection);
	result, err := collection.InsertMany(context.TODO(), newUsers, opts);
	if err != nil {
		fmt.Println("Error:", err)
	}
	jsonData, err := json.MarshalIndent(result, ""," ");
	if err != nil {
		fmt.Println("Marshal Error:", err)
	}

	fmt.Println("Result:", string(jsonData));
	fmt.Println()
	fmt.Println("Return type", reflect.TypeOf(result))
	fmt.Println();
}