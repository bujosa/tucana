package domain

import "gopkg.in/mgo.v2/bson"

type Planet struct {
    Id      bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
    Name    string        `json:"name,omitempty" bson:"name,omitempty"`
}
type IPlanetRespository interface {
	FindAll() ([]Planet, error)
	ById(bson.ObjectId) (*Planet, error)
}
