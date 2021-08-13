package app

import "net/http"

func(app *Application) IsPostMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			app.JSON(w, http.StatusMethodNotAllowed, &MessageResponse{http.StatusText(http.StatusMethodNotAllowed)})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func(app *Application) IsJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		contentType := r.Header.Get("Content-Type")
		if len(contentType) == 0 || contentType != "application/json" {
			app.JSON(w, http.StatusBadRequest,&MessageResponse{"Header missing: application/json"})
			return
		}
		next.ServeHTTP(w, r)
	})
}
