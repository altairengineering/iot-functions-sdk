package function

import (
	"io"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var input []byte

	if r.Body != nil {
		defer r.Body.Close()
		input, _ = io.ReadAll(r.Body)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(input)
}
