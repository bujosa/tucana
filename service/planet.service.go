package service

import (
	"saturn-golang/domain"

	"gopkg.in/mgo.v2/bson"
)


type IPlanetService interface {
	GetAllPlanets() ([]domain.Planet, error)
	GetPlanet(bson.ObjectId) (*domain.Planet, error)
}

type DefaultPlanetService struct {
	repo domain.IPlanetRespository
}

func (s DefaultPlanetService) GetAllPlanets() ([]domain.Planet, error) {
	return s.repo.FindAll()
}

func (s DefaultPlanetService) GetPlanet(id bson.ObjectId) (*domain.Planet, error) {
	return s.repo.ById(id)
}

// Create new planet
func NewPlanetService(repository domain.IPlanetRespository) DefaultPlanetService {
	return DefaultPlanetService{repo: repository}
}