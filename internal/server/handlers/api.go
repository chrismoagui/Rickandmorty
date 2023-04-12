package handlers

import (
	"gomysql/internal/data"
	"net/http"

	"github.com/go-chi/chi"
)

func New() http.Handler {
	r := chi.NewRouter()
	ch := &CharacterRouter{
		Repository: &data.CharacterRepository{
			Data: data.New(),
		},
	}

	lo := &LocationRouter{
		Repository: &data.LocationRepository{
			Data: data.New(),
		},
	}

	r.Mount("/location/", lo.Routes())
	r.Mount("/character", ch.Routes())

	return r
}
