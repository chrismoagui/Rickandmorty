package handlers

import (
	"encoding/json"
	"fmt"
	"gomysql/pkg/location"
	"gomysql/pkg/response"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type LocationRouter struct {
	Repository location.Repository
}

/*
	CREATE ARTICLE VIDEOGAME
*/

func (lo *LocationRouter) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var locationRepo location.Location

	err := json.NewDecoder(r.Body).Decode(&locationRepo)

	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
	}

	defer r.Body.Close()

	ctx := r.Context()
	err = lo.Repository.Create(ctx, &locationRepo)

	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
	}

	w.Header().Add("Location", fmt.Sprintf("%s%d", r.URL.String(), locationRepo.ID))
	response.JSON(w, r, http.StatusCreated, response.Map{"Location": locationRepo})

}

/*
	LIST Articles

*/

func (lo *LocationRouter) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	locations, err := lo.Repository.GetAll(ctx)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, locations)
}

/*
	DELETE ARTICLES
*/

func (lo *LocationRouter) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "idlocation")

	idlocation, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	err = lo.Repository.Delete(ctx, uint(idlocation))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"msg": "Delete location!"})
}

/*
	GET BY USER
*/

/*
	GET ONE ARTICLES
*/

func (lo *LocationRouter) GetOneHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "idlocation")

	idlocation, err := strconv.Atoi(idStr)
	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	location, err := lo.Repository.GetOne(ctx, uint(idlocation))
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	response.JSON(w, r, http.StatusOK, response.Map{"Location": location})
}

/*
	UPDATE ARTICLES
*/

func (lo *LocationRouter) UpdateHandler(w http.ResponseWriter, r *http.Request) {

	idStr := chi.URLParam(r, "idlocation")
	idlocation, err := strconv.Atoi(idStr)

	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	var locationRepo location.Location
	err = json.NewDecoder(r.Body).Decode(&locationRepo)

	if err != nil {
		response.HTTPError(w, r, http.StatusBadRequest, err.Error())
		//GOlang regresa los parametros de la funcion por defecto en el return, creo.
		return
	}

	defer r.Body.Close()

	ctx := r.Context()
	location, err := lo.Repository.Update(ctx, uint(idlocation), locationRepo)
	if err != nil {
		response.HTTPError(w, r, http.StatusNotFound, err.Error())
		return
	}

	//response.JSON(w, r, http.StatusOK, nil)
	response.JSON(w, r, http.StatusOK, response.Map{"Location": location})

}

//ROUTERS

func (lo *LocationRouter) Routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", lo.GetAllHandler)  //LIST
	r.Post("/", lo.CreateHandler) //CREATE

	r.Get("/{idlocation}/", lo.GetOneHandler) //DETAIL
	r.Put("/{idlocation}/", lo.UpdateHandler) //UPDATE

	r.Delete("/{idlocation}/", lo.DeleteHandler) //DELETE

	return r
}
