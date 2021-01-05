package domain

type Planet struct {
	id   string "json"
	name string
}

type IPlanetRespository interface {
	FindAll() ([]Planet, error)
}
