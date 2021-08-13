package app

import (
	"encoding/json"
	"net/http"
)

func(app *Application) JSON(w http.ResponseWriter, status int, body interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

