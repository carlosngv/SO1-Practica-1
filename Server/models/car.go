package models

import "gopkg.in/mgo.v2/bson"

type Car struct {
	Id		bson.ObjectId	`json:"id" bson:"_id"`
	Plate	string			`json:"plate" bson:"plate"`
	Brand	string			`json:"brand" bson:"brand"`
	Model	string			`json:"model" bson:"model"`
	Series	string			`json:"series" bson:"series"`
	Color	string			`json:"color" bson:"color"`
}
