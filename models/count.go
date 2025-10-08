package models

import (
	"fmt"
	"context"
    "go.mongodb.org/mongo-driver/bson"
	"reflect"
)

func CountAllDocuments() int64 {
	fmt.Print("---------------------------------\nCount All Documents\n---------------------------------\n")
	collection := mongoClient.Database(db).Collection(collection);
	count, err := collection.CountDocuments(context.TODO(), bson.M{});
	if err != nil {
		fmt.Println("Error", err);
	}

	fmt.Println("Count Result:", count);
	fmt.Println("Return Type:", reflect.TypeOf(count))
	fmt.Println();

	return count;
}

func CountDocumentsWithFilter(balance int) int64 {
	fmt.Print("---------------------------------\nCount With Filter\n---------------------------------\n")
	filter := bson.M{ "balance": bson.M{ "$gte": balance }}

	collection := mongoClient.Database(db).Collection(collection);
	count, err := collection.CountDocuments(context.TODO(), filter);

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Count Result:", count);
	fmt.Println("Return Type:", reflect.TypeOf(count))
	fmt.Println();

	return count;
}