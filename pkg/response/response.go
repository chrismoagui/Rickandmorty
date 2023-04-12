package response

import (
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

type Map map[string]interface{}

func JSON(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) error {
	if data == nil {
		//Allow CORS here By * or specific Origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(statusCode)
		return nil
	}

	j, err := json.Marshal(data)

	if err != nil {
		return err
	}

	//Allow CORS here By * or specific Origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(j)

	return nil

}

func HTTPError(w http.ResponseWriter, r *http.Request, statusCode int, message string) error {
	msg := ErrorMessage{
		Message: message,
	}
	return JSON(w, r, statusCode, msg)
}
