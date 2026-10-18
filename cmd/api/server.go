package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/zekielmp/Agent/internal/database"
)

type Application struct {
}

func Server(app *Application) error {
	dsn := os.Getenv("DB_DSN")
	db, err := database.ConnDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database:%s", err.Error())
	} else {
		log.Println("Database connected successfully")
	}

	defer db.Close()

	port := os.Getenv("PORT")
	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.routes(),
	}
	log.Printf("Server runnning on port :%s", port)
	return srv.ListenAndServe()
}
