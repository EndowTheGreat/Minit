package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"

	"gitlab.com/EndowTheGreat/minit/db"
	internalHTTP "gitlab.com/EndowTheGreat/minit/http"
)

func setup() error {
	handler := internalHTTP.NewHandler()
	handler.Router.Use(middleware.Logger, cors.AllowAll().Handler, httprate.LimitByIP(100, 1*time.Minute))
	handler.SetupRoutes()
	handler.DB = db.NewDB()
	defer handler.DB.Close()
	server := http.Server{
		Addr:    ":8080",
		Handler: handler.Router,
	}
	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("Attempting to instantiate Minit instance...")
	if err := setup(); err != nil {
		log.Fatal("Failed to instantiate Minit instance: ", err)
	}
}
