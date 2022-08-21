package models

import "gopkg.in/mgo.v2/bson"

type Log struct {
	Id		bson.ObjectId	`json:"id" bson:"_id"`
	Description	string			`json:"description" bson:"brand"`
	LogDate	string			`json:"logDate" bson:"logDate"`
}
