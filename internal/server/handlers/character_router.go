package handlers

import (
	"encoding/json"
	"fmt"
	"gomysql/pkg/character"
	"gomysql/pkg/response"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type CharacterRouter struct {
	Repository character.Repository
}

func (ch *CharacterRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var characterRepo character.Character

	err := json.NewDecoder(r.Body).Decode(&characterRepo)

	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = ch.Repository.Create(ctx, &characterRepo)

	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
	}

	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), characterRepo.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"character": characterRepo})

}

func (ch *CharacterRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	characters, err := ch.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, characters)
}

/*
	DELETE ARTICLES
*/

func (ch *CharacterRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "idcharacter")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = ch.Repository.Delete(ctx, uint(id))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"msg": "Delete character!"})
}

/*
	GET ONE ARTICLES
*/

func (ch *CharacterRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "idcharacter")

	idcharacter, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	article, err := ch.Repository.GetOne(ctx, uint(idcharacter))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"ArticleVideoGame": article})
}

/*
	UPDATE ARTICLES
*/

func (ch *CharacterRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "idcharacter")
	idcharacter, err := strconv.Atoi(idStr)

	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	var characterRepo character.Character
	err = json.NewDecoder(r.Body).Decode(&characterRepo)

	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		//GOlang regresa los parametros de la funcion por defecto en el return, creo.
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	character, err := ch.Repository.Update(ctx, uint(idcharacter), characterRepo)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	//response.JSON(w, r, http.StatusOK, nil)
	response.JSON(w, r, http.StatusOK, response.Map{"Character": character})

}

//ROUTERS

func (ch *CharacterRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", ch.GetAllHandler)  //LIST
	r.Post("/", ch.CreateHandler) //CREATE

	r.Get("/{idcharacter}/", ch.GetOneHandler) //DETAIL
	r.Put("/{idcharacter}/", ch.UpdateHandler) //UPDATE

	r.Delete("/{idcharacter}/", ch.DeleteHandler) //DELETE

	return r
}
