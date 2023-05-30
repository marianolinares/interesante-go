package healthcheck

import (
	"net/http"
)

func CreateHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All is OK!"))
	}
}
