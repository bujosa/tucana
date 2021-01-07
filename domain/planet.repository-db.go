package domain

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)

type PlanetRepositoryDb struct {
   collection *mgo.Collection
}

var collection = getSession().DB("saturn").C("planets")

func getSession() *mgo.Session {
   err := godotenv.Load()
   if err != nil {
     log.Fatal("Error loading .env file")
   }
   session, err := mgo.Dial(os.Getenv("MONGO_URL"))
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