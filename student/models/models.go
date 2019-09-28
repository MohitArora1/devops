package models

import "gopkg.in/mgo.v2/bson"

//Student model
type Student struct {
	ID     bson.ObjectId `json:"id" bson:"_id"`
	Name   string        `json:"name"`
	Age    int           `json:"age"`
	Class  string        `json:"class"`
	RollNo int           `json:"rollno"`
}
