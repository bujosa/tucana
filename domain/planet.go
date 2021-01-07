package domain

import "gopkg.in/mgo.v2/bson"

type Planet struct {
	_id  bson.ObjectId    `json:"id,omitempty"`
	name string           `json:"name,omitempty"`
}
type IPlanetRespository interface {
	FindAll() ([]Planet, error)
	ById(bson.ObjectId) (*Planet, error)
}
