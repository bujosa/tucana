package domain

import (
	"log"

	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)
type PlanetRepositoryDb struct {
}

var collection = getSession().DB("saturn").C("planets")

func getSession() *mgo.Session {
   session, err := mgo.Dial("mongodb://localhost")
   if err != nil {
	 	panic(err)
   }
   return session
}

func (d PlanetRepositoryDb) FindAll() ([]Planet, error) {
   var results []Planet
   err := collection.Find(nil).All(&results)
   if err != nil {
	   log.Fatal(err)
   }
   return results, nil
}

func (d PlanetRepositoryDb) ById(id string) (*Planet, error){
   result := Planet{}
   err:= collection.FindId(id).One(&result)

   if err != nil{
      log.Fatal(err)
   }
   return &result, nil
}

func NewPlanetRepositoryDb() PlanetRepositoryDb {
	
}