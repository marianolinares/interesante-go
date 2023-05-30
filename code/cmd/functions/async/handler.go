package async

import (
	"encoding/json"
	"fmt"
	"marian.com/interesante-go/code/internal"
	"marian.com/interesante-go/code/internal/registerStock"
	"net/http"
	"strconv"
)

type InvokeRequest struct {
	Data     map[string]json.RawMessage
	Metadata map[string]interface{}
}

type Message struct {
	Name   string `json:"name"`
	Legs   int    `json:"legs"`
	Colour string `json:"colour,omitempty"`
}

func CreateHandler(useCase registerStock.UseCase) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		body := parseBody(r)
		mappedEntity := internal.NewEntity(body.Legs, body.Name)

		useCase.Execute(mappedEntity)

		fmt.Println(body)
		fmt.Println(mappedEntity)

		w.WriteHeader(http.StatusOK)
	}
}

func parseBody(r *http.Request) (message Message) {
	var invokeRequest InvokeRequest
	json.NewDecoder(r.Body).Decode(&invokeRequest)

	var parsedMessage string
	json.Unmarshal(invokeRequest.Data["async"], &parsedMessage)

	data, _ := strconv.Unquote(parsedMessage)

	json.Unmarshal([]byte(data), &message)
	return
}
