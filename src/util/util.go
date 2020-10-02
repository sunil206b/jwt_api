package util

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/sunil206b/jwt_api/src/model"
	"golang.org/x/crypto/bcrypt"
)

// EncryptPassword function will encrypt the password from plain text to hash
func EncryptPassword(plainPass string) (string, error) {
	password := []byte(plainPass)
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Error while encoding password: %v", err)
	}
	return string(hash), nil
}

// ComparePassword function will decrypt the encrypted password to palin text
func ComparePassword(hashedPass string, plainPass string) (bool, error) {
	password := []byte(plainPass)
	encryptPassword := []byte(hashedPass)

	err := bcrypt.CompareHashAndPassword(encryptPassword, password)
	if err != nil {
		return false, fmt.Errorf("Not a valid password: %v", err)
	}
	return true, nil
}

// GenerateToken function will generate the signed token based on the given sign
func GenerateToken(secret string, userID int, email string) (string, error) {
	claims := model.MyClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Issuer:    email,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	ss, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("Error while creating signed token: %v", err)
	}
	return ss, nil
}

// VerifyToken function will take the signed token from the client and returns the
func VerifyToken(tokenString string, secret string) (*model.MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	claims, ok := token.Claims.(*model.MyClaims)
	if !ok && !token.Valid {
		return nil, fmt.Errorf("Not a valid token: %v", err)
	}
	return claims, nil
}
