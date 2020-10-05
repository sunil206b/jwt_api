package service

import (
	"context"
	"fmt"
	"github.com/sunil206b/jwt_api/src/model"
	"github.com/sunil206b/jwt_api/src/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strings"
)

type Service struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewService(coll *mongo.Collection, ctx context.Context) *Service {
	return &Service{
		collection: coll,
		ctx:        ctx,
	}
}

func (s *Service) CreateUser(user *model.User) *util.RestErr {
	res, err := s.collection.InsertOne(s.ctx, user)
	if err != nil {
		if strings.Contains(err.Error(), user.Email) {
			log.Printf("error while creating new user unique constraint %s\n", user.Email)
			return util.NewBadRequest(fmt.Sprintf("email already exists %s", user.Email))
		}
		log.Printf("error while creating new user %v\n", err)
		return util.NewInternalServerError("error while creating new user")
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		util.NewInternalServerError("error while getting user id from mongodb")
	}
	user.ID = oid
	return nil
}

func (s *Service) GetUser(loginUse *model.LoginUser) (*model.User, *util.RestErr) {
	filter := bson.M{"email": loginUse.Email}
	user := new(model.User)
	err := s.collection.FindOne(s.ctx, filter).Decode(user)
	if err != nil {
		return nil, util.NewNotFoundError("user not found with the give id")
	}
	return user, nil
}
