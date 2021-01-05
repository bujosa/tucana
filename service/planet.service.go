package service

import "saturn-golang/domain"

type IPlanetService interface {
	GetAllPlanets() ([]domain.Planet, error)
}

type DefaultPlanetService struct {
	repo domain.IPlanetRespository
}

func (s DefaultPlanetService) GetAllPlanets() ([]domain.Planet, error) {
	return s.repo.FindAll()
}

func NewPlanetService(repository domain.IPlanetRespository) DefaultPlanetService {
	return DefaultPlanetService{repo: repository}
}