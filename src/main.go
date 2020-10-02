package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sunil206b/jwt_api/src/controllers"
	"github.com/sunil206b/jwt_api/src/driver"
)

func main() {
	client, err := driver.MongoDB()
	if err != nil {
		log.Fatalln(err)
	}

	controller := controllers.NewController(client)
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Index).Methods(http.MethodGet)
	r.HandleFunc("/login", controller.Login).Methods(http.MethodPost)
	r.HandleFunc("/signup", controller.SignUp).Methods(http.MethodPost)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	if err = srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
