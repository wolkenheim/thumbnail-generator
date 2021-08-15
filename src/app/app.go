package app

import "net/http"

type ApplicationInterface interface {
	JSON(w http.ResponseWriter, status int, body interface{})
	Liveness(w http.ResponseWriter, req *http.Request)
	IsPostMiddleware(next http.Handler) http.Handler
	IsJSONMiddleware(next http.Handler) http.Handler
}

type Application struct {
}

func(app *Application) Liveness(w http.ResponseWriter, req *http.Request) {
	app.JSON(w, http.StatusOK, `{"message":"all good"}`)
}

func NewApplication() *Application{
	return &Application{}
}
