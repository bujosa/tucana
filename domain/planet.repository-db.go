package domain

import (
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)
type PlanetRepositoryDb struct {
}

func getSession() *mgo.Session {
   session, err := mgo.Dial("mongodb://localhost")

   if err != nil {
	 	panic(err)
   }

   return session
}

func (d PlanetRepositoryDb) FindAll() ([]Planet, error) {

}