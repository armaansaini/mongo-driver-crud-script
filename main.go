package main

import (
	"go-test-mongo/models"
	"fmt"
)

/*

1. Try duplicate id in insert bulk run to see functionality of unordered bulk update.
2. Try to timeout the query with findWithTimeout



*/



func main(){

	fmt.Println("\n")

	models.ConnectDatabase()

	fmt.Println("Truncate all data....")
	models.DeleteAll()

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\nInsert Operations\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	insertedId := models.InsertOne();
	models.InsertMany();
	models.InsertManyUnordered();

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\nCount Operations\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	models.CountAllDocuments();
	models.CountDocumentsWithFilter(3000); // add balance as filter

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\nRead Operations\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	models.FindOne() // find first document
	models.FindById(insertedId) // find by id
	models.FindByName("test1")
	models.FindAllWithFilter(5000) // use balance as filter
	models.FindAll();
	models.IterRead();
	models.FindAllWithOptions()
	models.FindWithTimeout() // mongo.CommandError
	models.FindWithContextTimeout() //context.deadlineExceededError

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\nUpdate Operations\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	models.UpdateOne("test10", 5700);
	models.UpdateOneById(insertedId, 250);
	models.UpdateOneWithUpsert("test200", 9999)
	models.UpdateMany(50, 330); // age >= 50, update balance
	models.UpdateManyWithUpsert(100, 9049) // age >= 100, update balance

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\nDelete Operations\n~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	models.DeleteOne("test2")
	models.DeleteOneById(insertedId)
	models.DeleteMany(6000) // balance >= 6000


}