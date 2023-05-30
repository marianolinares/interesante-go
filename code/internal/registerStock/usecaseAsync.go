package registerStock

import (
	"context"
	"log"
	"marian.com/interesante-go/code/internal"
)

type UseCase interface {
	Execute(ctx context.Context, entity internal.Entity) (string, error)
}

type usecase struct {
	repo internal.EntityRepo
}

func NewUseCase(r internal.EntityRepo) UseCase {
	return &usecase{r}
}

func (u *usecase) Execute(ctx context.Context, entity internal.Entity) (string, error) {
	data, err := u.repo.GetEntities()

	if err != nil {
		return "", err
	}

	log.Println(data)
	u.repo.SaveEntity(entity)
	log.Println("Entity saved", entity)

	return data[0].Name, nil
}
