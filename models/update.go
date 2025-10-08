package models

import (
	"fmt"
	"encoding/json"
	"context"
	"reflect"
	"go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)


// update by field
func UpdateOne(name string, balance int) error {
	fmt.Print("---------------------------------\nUpdate One - Update by field\n---------------------------------\n")
	filter := bson.M{ "name": name}
	update := bson.M{"$set": bson.M{"balance": balance}}

	collection := mongoClient.Database(db).Collection(collection);
	result, err := collection.UpdateOne(context.TODO(), filter, update);
	if err != nil {
		fmt.Println("Error:", err)
	}
	jsonData, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		fmt.Println("Marshal Error:", err)
	}

	fmt.Println("Result:", string(jsonData));
	fmt.Println()
	fmt.Println("Return type:", reflect.TypeOf(result))
	fmt.Println("MatchedCount Type:", reflect.TypeOf(result.MatchedCount))
	fmt.Println("ModifiedCount Type:", reflect.TypeOf(result.ModifiedCount))
	fmt.Println("UpsertedCount Type:", reflect.TypeOf(result.UpsertedCount))
	fmt.Println("UpsertedID Type:", reflect.TypeOf(result.UpsertedID))
	fmt.Println()
	return nil;
}

// update by id
func UpdateOneById(userId string, balance int) error {
	fmt.Print("---------------------------------\nUpdate One By Id\n---------------------------------\n")
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		fmt.Println("ObjectID Error:", err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"balance": balance}}

	collection := mongoClient.Database(db).Collection(collection);
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println("Error:", err)
	}
	jsonData, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		fmt.Println("Marshal Error:", err)
	}

	fmt.Println("Result:", string(jsonData));
	fmt.Println()
	fmt.Println("Return type:", reflect.TypeOf(result))
	fmt.Println("MatchedCount Type:", reflect.TypeOf(result.MatchedCount))
	fmt.Println("ModifiedCount Type:", reflect.TypeOf(result.ModifiedCount))
	fmt.Println("UpsertedCount Type:", reflect.TypeOf(result.UpsertedCount))
	fmt.Println("UpsertedID Type:", reflect.TypeOf(result.UpsertedID))
	fmt.Println()
	return nil;
}

func UpdateOneWithUpsert(name string, balance int) error {
	fmt.Print("---------------------------------\nUpdate One - Upsert\n---------------------------------\n")
	filter := bson.M{ "name": name }
	update := bson.M{ "$set": bson.M{ "balance": balance }}
	opts := options.Update().SetUpsert(true)

	collection := mongoClient.Database(db).Collection(collection);
	result, err := collection.UpdateOne(context.TODO(), filter, update, opts);
	if err != nil {
		fmt.Println("Error:", err)
	}
	jsonData, err := json.MarshalIndent(result, "", " ");
	if err != nil {
		fmt.Println("Marshal Error:", err)
	}

	fmt.Println("Result:", string(jsonData));
	fmt.Println()
	fmt.Println("Return type:", reflect.TypeOf(result))
	fmt.Println("MatchedCount Type:", reflect.TypeOf(result.MatchedCount))
	fmt.Println("ModifiedCount Type:", reflect.TypeOf(result.ModifiedCount))
	fmt.Println("UpsertedCount Type:", reflect.TypeOf(result.UpsertedCount))
	fmt.Println("UpsertedID Type:", reflect.TypeOf(result.UpsertedID))
	fmt.Println()
	return err;
}

func UpdateMany(age int, balance int) error {
	fmt.Print("---------------------------------\nUpdate Many\n---------------------------------\n")
	filter := bson.M{ "name": bson.M{ "$gte": age} }
	update := bson.M{ "$set" : bson.M{ "balance": balance }}

	collection := mongoClient.Database(db).Collection(collection);
	result, err := collection.UpdateMany(context.TODO(), filter, update);
	if err != nil {
		fmt.Println("Error:", err)
	}
	jsonData, err := json.MarshalIndent(result, "" , " ");
	if err != nil {
		fmt.Println("Marshal Error:", err)
	}

	fmt.Println("Result:", string(jsonData));
	fmt.Println()
	fmt.Println("Return type:", reflect.TypeOf(result))
	fmt.Println("MatchedCount Type:", reflect.TypeOf(result.MatchedCount))
	fmt.Println("ModifiedCount Type:", reflect.TypeOf(result.ModifiedCount))
	fmt.Println("UpsertedCount Type:", reflect.TypeOf(result.UpsertedCount))
	fmt.Println("UpsertedID Type:", reflect.TypeOf(result.UpsertedID))
	fmt.Println()
	return nil
}

func UpdateManyWithUpsert(age int, balance int) error {
	fmt.Print("---------------------------------\nUpdate Many - Upsert\n---------------------------------\n")
	filter := bson.M{ "age": bson.M{ "$gte": age } }
	update := bson.M{ "$set" : bson.M{ "balance": balance }}
	options := options.Update().SetUpsert(true)

	collection := mongoClient.Database(db).Collection(collection);
	result, err := collection.UpdateMany(context.TODO(), filter, update, options);
	if err != nil {
		fmt.Println("Error:", err)
	}
	jsonData, err := json.MarshalIndent(result, "" , " ");
	if err != nil {
		fmt.Println("Marshal Error:", err)
	}

	fmt.Println("Result:", string(jsonData));
	fmt.Println()
	fmt.Println("Return type:", reflect.TypeOf(result))
	fmt.Println("MatchedCount Type:", reflect.TypeOf(result.MatchedCount))
	fmt.Println("ModifiedCount Type:", reflect.TypeOf(result.ModifiedCount))
	fmt.Println("UpsertedCount Type:", reflect.TypeOf(result.UpsertedCount))
	fmt.Println("UpsertedID Type:", reflect.TypeOf(result.UpsertedID))
	fmt.Println()

	return nil
}