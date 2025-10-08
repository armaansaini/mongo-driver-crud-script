package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/bson"
)

type User struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	Name string			  `json:"name"`
	Age  int			  `json:"age" bson:"age"`
	Balance int		  	  `json:"balance"`
}

var newUser = bson.M{ "name": "test1", "age": 10, "balance": 2500}

var users = []bson.M{
	{ "name": "test2", "age": 20, "balance": 10000},
	{ "name": "test3", "age": 30, "balance": 5000},
	{ "name": "test4", "age": 15, "balance": 7733},
}

