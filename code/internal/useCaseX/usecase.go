package useCaseX

import (
	"log"
	"marian.com/interesante-go/code/internal"
)

func FetchData(repository internal.EntityRepo) {
	data, err := repository.GetEntities()

	if err != nil {
		log.Fatal(err)
	}

	log.Println(data)
}
