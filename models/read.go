package models

import (
	"fmt"
	"encoding/json"
	"context"
	"reflect"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Find One
func FindOne() User{
	fmt.Print("---------------------------------\nFindOne() - No filter \n---------------------------------\n")
	var result User;

	collection := mongoClient.Database(db).Collection(collection);
	err := collection.FindOne(context.TODO(), bson.M{}).Decode(&result);
	if err != nil {
		fmt.Println("Error:", err)
	}
	jsonData, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		fmt.Println("Marshal Error:", err)
	}

	fmt.Println("Result:",string(jsonData))
	fmt.Println("Return Type:", reflect.TypeOf(result));
	fmt.Println("\n");
	return result

}

// Find All
func FindAll() []User {
	fmt.Print("---------------------------------\nFind() - No filter \n---------------------------------\n")
	var results []User;

	collection := mongoClient.Database(db).Collection(collection);
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println("Error:", err)
	}
	cursor.All(context.TODO(), &results);
	if err != nil {
		fmt.Println("Marshal Error:", err)
	}

	jsonData, err := json.MarshalIndent(results, "" ," ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Result:",string(jsonData))
	fmt.Println("Return Type:", reflect.TypeOf(results));
	fmt.Println("\n");
	return results;
}

// find one by name
func FindByName(username string) User {
	fmt.Print("---------------------------------\nFindOne() - With Filter \n---------------------------------\n")
	var result User
	filter := bson.M{ "name": username }
	
	collection := mongoClient.Database(db).Collection(collection)
	err := collection.FindOne(context.TODO(), filter).Decode(&result);
	if err != nil {
		fmt.Println("Error:", err)
	}
	jsonData, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		fmt.Println("Marshal Error:", err)
	}

	fmt.Println("Result:",string(jsonData))
	fmt.Println("Return Type:", reflect.TypeOf(result));
	fmt.Println();
	return result
}

// find one by id
func FindById(userId string) User {
	fmt.Print("---------------------------------\nFindOne() - With Id \n---------------------------------\n")
	id, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		fmt.Println("ObjectId Error:", err)
	}

	filter := bson.M{ "_id": id }

	collection := mongoClient.Database(db).Collection(collection);
	var result User

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println("Error:", err)
	}
	jsonData, err := json.MarshalIndent(result, ""," ")
	if err != nil {
		fmt.Println("Marshal Error:", err)
	}

	fmt.Println("Result:",string(jsonData))
	fmt.Println("Return Type:", reflect.TypeOf(result));
	// fmt.Println("id", result.ID, "return type", reflect.TypeOf(result.ID))
	// fmt.Println("age", result.Age, "return type", reflect.TypeOf(result.Age))
	fmt.Println("\n");
	return result
}

// find all with filter
func FindAllWithFilter(balance int) []User {
	fmt.Print("---------------------------------\nFindAll() - With filter \n---------------------------------\n")
	var results []User;

	filter := bson.M{ "balance": bson.M{ "$gte": balance }}

	collection := mongoClient.Database(db).Collection(collection);
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = cursor.All(context.TODO(), &results);
	if err != nil {
		fmt.Println("Error:", err)
	}
	jsonData, err := json.MarshalIndent(results, "" , " ")
	if err != nil {
		fmt.Println("Marshal Error:", err)
	}

	fmt.Println("Result:",string(jsonData))
	fmt.Println("Return Type:", reflect.TypeOf(results));
	// fmt.Println("id", result.ID, "return type", reflect.TypeOf(result.ID))
	// fmt.Println("age", result.Age, "return type", reflect.TypeOf(result.Age))
	fmt.Println("\n");
	return results;
}

func IterRead(){
	fmt.Print("---------------------------------\nFind() - Batch Cursor read \n---------------------------------\n")

	findOptions := options.Find().SetBatchSize(1)
	cursor, err := mongoClient.Database(db).Collection(collection).Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		fmt.Println("Error", err)
	}
	defer cursor.Close(context.TODO())

	for true {
		var result User;
		success := cursor.Next(context.TODO())
		err := cursor.Decode(&result); 	
		if !success {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
		}
		jsonData, err := json.MarshalIndent(result, "", " ");
		if err != nil {
			fmt.Println("Marshal Error:", err);
		}
		fmt.Println("Result:",string(jsonData))
		fmt.Println("Return Type:", reflect.TypeOf(result));
	}
	fmt.Println("\n")
}

func FindAllWithOptions(){
	fmt.Print("---------------------------------\nFind() - With Sort(name,desc) and Limit(3) \n---------------------------------\n")
	opts := options.Find().
        SetSort(bson.D{{Key: "name", Value: -1}}).
		SetLimit(3)

	collection := mongoClient.Database(db).Collection(collection);
	cursor, err := collection.Find(context.TODO(), bson.M{}, opts);
	if err != nil {
		fmt.Println("Error:", err)
	}
	var results []User

	err = cursor.All(context.TODO(), &results)
	if err != nil {
		fmt.Println("Error:", err)
	}
	jsonData, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		fmt.Println("Marshal Error:", err)
	}
	
	fmt.Println("Result:",string(jsonData))
	fmt.Println("Return Type:", reflect.TypeOf(results));
	fmt.Println("\n");
}

func FindWithTimeout(){
	fmt.Print("---------------------------------\nFind() - SetMaxTime() \n---------------------------------\n")
	opts := options.Find().
				SetSort(bson.D{{ Key: "Name", Value: -1 }}).
				SetLimit(1).
				SetMaxTime(1 * time.Nanosecond)

	filter := bson.M{};
	// uncomment to make it fail
	// filter := bson.M{"$where": "function() { sleep(50); return true; }"} // MongoDB server-side sleep

	collection := mongoClient.Database(db).Collection(collection)
	var result []User
	cursor, err := collection.Find(context.TODO(), filter, opts);
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	cursor.All(context.TODO(), &result)
	jsonData, err := json.MarshalIndent(result, "", " ");
	if err != nil {
		fmt.Println("Marshal Error:", err)
	}

	fmt.Println("Result:", string(jsonData))
	fmt.Println("Return Type:", reflect.TypeOf(result));
	fmt.Println("\n");
}

func FindWithContextTimeout(){
	fmt.Print("---------------------------------\nFind() - Context Timeout \n---------------------------------\n")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	opts := options.Find().
				SetSort(bson.D{{ Key: "Name", Value: -1 }}).
				SetLimit(1)

	filter := bson.M{}
	// uncomment to make it fail
	// filter := bson.M{"$where": "function() { sleep(50); return true; }"} // MongoDB server-side sleep

	collection := mongoClient.Database(db).Collection(collection)
	cursor, err := collection.Find(ctx, filter, opts);
	if err != nil {
		fmt.Println("Error:", err)
		return;
	}
	var results []User
	cursor.All(context.TODO(), &results)

	jsonData, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		fmt.Println("Marshal Error:", err)
	}

	fmt.Println("Result:", string(jsonData))
	fmt.Println("Return Type:", reflect.TypeOf(results));
	fmt.Println("\n");
}