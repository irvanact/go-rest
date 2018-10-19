package model

import "gopkg.in/mgo.v2/bson"

type Account struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Email     string        `bson:"email" json:"email"`
	Firstname string        `bson:"firstname" json:"firstname"`
	Lastname  string        `bson:"lastname" json:"lastname"`
}
