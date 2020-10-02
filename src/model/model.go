package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User will hold the user information when creating new user
type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    int                `json:"userId" bson:"user_id"`
	FirstName string             `json:"firstName" bson:"first_name"`
	LastName  string             `json:"lastName" bson:"last_name"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updated_at"`
}

// LoginUser will hold the information when user try to login
type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// MyClaims will hold the claims information
type MyClaims struct {
	SessionID string
	jwt.StandardClaims
}
