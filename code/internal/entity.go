package internal

type Entity struct {
	Price int
	Name  string
}

func NewEntity(price int, name string) (Entity, error) {
	e := Entity{
		Price: price,
		Name:  name,
	}
	return e, nil
}

func (e Entity) PriceCountry() string {
	return string(e.Price) + " " + e.Name
}

type EntityRepo interface {
	GetEntities() ([]Entity, error)
	SaveEntity(e Entity)
}
