package models

import (
	"fmt"
	"reflect"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/bson"
	"context"
)

// delete one by field
func DeleteOne(name string) error {
	fmt.Print("---------------------------------\nDelete One by field\n---------------------------------\n")
	filter := bson.M{ "name": name }
	
	collection := mongoClient.Database(db).Collection(collection);
	result, err := collection.DeleteOne(context.TODO(), filter);
	if err != nil {
		fmt.Println("Error:", err)
	}
	jsonData, err := json.MarshalIndent(result, ""," ")
	if err != nil {
		fmt.Println("Marshal Error:", err)
	}

	fmt.Println("Result:", string(jsonData));
	fmt.Println();
	fmt.Println("Return Type:", reflect.TypeOf(result))
	fmt.Println("DeletedCount Type:", reflect.TypeOf(result.DeletedCount))
	fmt.Println();
	return nil;
}

// delete one by id
func DeleteOneById(id string) {
	fmt.Print("---------------------------------\nDelete One by ID\n---------------------------------\n")
	pId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{ "_id": pId }

	collection := mongoClient.Database(db).Collection(collection);
	result, err := collection.DeleteOne(context.TODO(), filter);
	if err != nil {
		fmt.Println("Error:", err)
	}
	jsonData, err := json.MarshalIndent(result, "", " ");
	if err != nil {
		fmt.Println("Marshal error:", err)
	}

	fmt.Println("Result:", string(jsonData));
	fmt.Println();
	fmt.Println("Return Type:", reflect.TypeOf(result))
	fmt.Println("DeletedCount Type:", reflect.TypeOf(result.DeletedCount))
	fmt.Println();
}


// delete many
func DeleteMany(balance int) error {
	fmt.Print("---------------------------------\nDelete Many\n---------------------------------\n")
	filter := bson.M{ "balance": bson.M{ "$gte": balance } }

	collection := mongoClient.Database(db).Collection(collection);
	result, err := collection.DeleteMany(context.TODO(), filter);
	if err != nil {
		fmt.Println("Error:", err)
	}
	jsonData, err := json.MarshalIndent(result, "", " ");
	if err != nil {
		fmt.Println("Marshal Error:", err)
	}

	fmt.Println("Result:", string(jsonData));
	fmt.Println();
	fmt.Println("Return Type:", reflect.TypeOf(result))
	fmt.Println("DeletedCount Type:", reflect.TypeOf(result.DeletedCount))
	fmt.Println();
	return nil;
}

func DeleteAll() error {
	collection := mongoClient.Database(db).Collection(collection);
	_, err := collection.DeleteMany(context.TODO(), bson.M{});
	// _, err := json.MarshalIndent(result, "", " ");
	if err != nil {
		fmt.Println("Error:", err)
	}

	// fmt.Println("result", result);
	// fmt.Println("Result:", string(jsonData));
	// fmt.Println()
	// fmt.Println("Return type:", reflect.TypeOf(result));
	// fmt.Println("Result.DeletedCount", result.DeletedCount, "Type", reflect.TypeOf(result.DeletedCount))

	return nil;
}