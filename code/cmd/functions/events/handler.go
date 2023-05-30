package events

import (
	"fmt"
	"marian.com/interesante-go/code/internal"
	"net/http"
)

func CreateHandler(entityRepo internal.EntityRepo) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		message := "This HTTP triggered function executed successfully. Pass a name in the query string for a personalized response.\n"
		name := r.URL.Query().Get("name")
		if name != "" {
			message = fmt.Sprintf("Hello, %s. This HTTP triggered function executed successfully.\n", name)

			entityRepo.SaveEntity(internal.NewEntity(10, name))
		}
		fmt.Fprint(w, message)
	}
}
