package domain

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
)

type PlanetRepositoryDb struct {
   collection *mgo.Collection
}

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
   err := d.collection.Find(nil).All(&results)
   fmt.Println("resultados:", results[1].name)
   if err != nil {
	   log.Fatal(err)
   }
   return results, nil
}

func (d PlanetRepositoryDb) ByName(name string) (*Planet, error){
   result := Planet{}
   err:= d.collection.Find({name:nam}).One(&result)

   if err != nil{
      log.Fatal(err)
   }
   return &result, nil
}

func NewPlanetRepositoryDb() PlanetRepositoryDb {
   collection := getSession().DB("saturn").C("planets")
   return PlanetRepositoryDb{collection: collection}
}