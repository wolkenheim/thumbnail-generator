package app

import "net/http"

type Application struct {
}

func(app *Application) Liveness(w http.ResponseWriter, req *http.Request) {
	app.JSON(w, http.StatusOK, `{"message":"all good"}`)
}
