package main

import (
	"encoding/json"
	"net/http"
)

func (app *Application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /ping", app.healthCheck)

	return mux
}

func (app *Application) healthCheck(w http.ResponseWriter, r *http.Request) {

	res := map[string]string{
		"Author":     "Giwa Ezekiel",
		"Github_url": "https://github.com/zekielmp",
		"Message":    "Server is live...🤧",
	}

	w.WriteHeader(http.StatusOK)
	data, _ := json.MarshalIndent(res, "", "\t")
	_, err := w.Write(data)
	if err != nil {
		http.Error(w, "An error occurred, please try again", http.StatusBadRequest)
		return
	}

}
