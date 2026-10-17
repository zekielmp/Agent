package main

import (
	"log"
	"net/http"
	"os"

	"github.com/zekielmp/Agent/internal/database"
)

type Application struct {
}

func Server(app *Application) error {
	db, err := database.ConnDB(os.Getenv("DB_DSN"))
	if err != nil {
		log.Println("Failed to connect to database:", err)
	}

	defer db.Close()

	srv := http.Server{
		Addr:    ":8080",
		Handler: app.SetupRoutes(),
	}

	return srv.ListenAndServe()
}
