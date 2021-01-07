package domain

type Planet struct {
	_id  string
	name string
}

type IPlanetRespository interface {
	FindAll() ([]Planet, error)
	ById(string) (*Planet, error)
}
