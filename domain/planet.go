package domain

type Planet struct {
	id   string
	name string
}

type IPlanetRespository interface {
	FindAll() ([]Planet, error)
	ById(string) (*Planet, error)
}
