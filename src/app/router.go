package app

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/sunil206b/jwt_api/src/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

func NewRouter(collection *mongo.Collection, ctx context.Context) *http.Server {

	controller := controllers.NewController(collection, ctx)
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Index).Methods(http.MethodGet)
	r.HandleFunc("/login", controller.Login).Methods(http.MethodPost)
	r.HandleFunc("/signup", controller.SignUp).Methods(http.MethodPost)
	r.HandleFunc("/token", controller.Token).Methods(http.MethodGet)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return srv
}
