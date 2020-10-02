package controllers

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

// Controller will hold driver information
type Controller struct {
	client *mongo.Client
}

// NewController function will return the controller
func NewController(client *mongo.Client) *Controller {
	return &Controller{
		client: client,
	}
}

// Index method will allow the user to create new user
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {

}

// Login method will receive the user credentials and validate them
// and creates the JWT token
func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) SignUp(w http.ResponseWriter, r *http.Request) {

}