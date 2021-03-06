package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sunil206b/jwt_api/src/model"
	"github.com/sunil206b/jwt_api/src/service"
	"github.com/sunil206b/jwt_api/src/util"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

// Controller will hold driver information
type Controller struct {
	service *service.Service
}

// NewController function will return the controller
func NewController(collection *mongo.Collection, ctx context.Context) *Controller {
	return &Controller{
		service: service.NewService(collection, ctx),
	}
}

// Index method will allow the user to create new user
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {

}

// Login method will receive the user credentials and validate them
// and creates the JWT token
func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	loginUser := new(model.LoginUser)
	err := json.NewDecoder(r.Body).Decode(loginUser)
	if err != nil {
		log.Printf("error while trying to marshal user %v\n", err)
		errMsg := util.NewBadRequest("not a valid login credentials")
		util.ResponseError(w, errMsg)
		return
	}
	user, errMsg := c.service.GetUser(loginUser)
	if errMsg != nil {
		log.Println(errMsg)
		util.ResponseError(w, errMsg)
		return
	}
	isValid, err := util.ComparePassword(user.Password, loginUser.Password)
	if err != nil || !isValid {
		log.Println("not valid password", err)
		errMsg := util.NewAuthenticationError("not a valid password")
		util.ResponseError(w, errMsg)
		return
	}
	sessionId := uuid.New().String()
	token, err := util.GenerateToken("secret", sessionId, user.Email)
	if err != nil {
		log.Println(err)
		errMsg := util.NewInternalServerError("error while generating token")
		util.ResponseError(w, errMsg)
		return
	}

	cookie := http.Cookie{
		Name:     "session-id",
		Value:    token,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	util.ResponseJson(w, http.StatusOK, user)
}

func (c *Controller) SignUp(w http.ResponseWriter, r *http.Request) {
	user := new(model.User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		log.Printf("error while trying to marshal user %v\n", err)
		errMsg := util.NewBadRequest("not a valid User")
		util.ResponseError(w, errMsg)
		return
	}
	user.ConstructUser()
	hashPass, err := util.EncryptPassword(user.Password)
	if err != nil {
		log.Printf("error while trying encrypt the user password %v\n", err)
		errMsg := util.NewInternalServerError(err.Error())
		util.ResponseError(w, errMsg)
		return
	}
	user.Password = hashPass
	if errMsg := c.service.CreateUser(user); errMsg != nil {
		log.Println(errMsg)
		util.ResponseError(w, errMsg)
		return
	}
	util.ResponseJson(w, http.StatusOK, user)
}

func (c *Controller) Token(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session-id")
	if err != nil {
		log.Println(err)
		errMsg := util.NewAuthenticationError("not a valid token")
		util.ResponseError(w, errMsg)
		return
	}
	fmt.Println(cookie.Value)
	if cookie.Value != "" {
		myToken, err := util.VerifyToken(cookie.Value, "secret")
		if err != nil {
			log.Println(err)
			errMsg := util.NewAuthenticationError("not a valid token")
			util.ResponseError(w, errMsg)
			return
		}
		util.ResponseJson(w, http.StatusOK, myToken)
		return
	}
	errMsg := util.NewAuthenticationError("not a valid token")
	util.ResponseError(w, errMsg)
}
