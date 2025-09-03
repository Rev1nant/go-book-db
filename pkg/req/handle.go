package req

import (
	"net/http"
	"strconv"

	"github.com/Rev1nant/go-book-db/pkg/res"
)

func Body[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		res.Json(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	if err := IsValid(body); err != nil {
		res.Json(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}

	return &body, nil
}

func GetId(w *http.ResponseWriter, r *http.Request) (int, error) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		res.Json(*w, "Invalid by", http.StatusBadRequest)
		return 0, err
	}

	return id, nil
}
