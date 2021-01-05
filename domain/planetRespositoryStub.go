package domain

type PlanetRespositoryStub struct {
	planets []Planet
}

func (s PlanetRespositoryStub) FindAll() ([]Planet, error) {
	return s.planets, nil
}

func NewPlanetRepositoryStub() PlanetRespositoryStub {
	planets := []Planet{
		{id: "1", name: "Saturn"},
		{id: "2", name: "Mars"},
	}

	return PlanetRespositoryStub{planets: planets}
}