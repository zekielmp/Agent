package main

import (
	"net/http"

	resp "github.com/zekielmp/Agent/internal/helpers"
)

func (app *Application) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", app.healthCheck)

	return mux
}

func (app *Application) healthCheck(w http.ResponseWriter, r *http.Request) {
	_ = resp.WriteJSON(w, http.StatusOK, "Server is live")
}
