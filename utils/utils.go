package utils

import (
  "time"

  "github.com/RrNn/detector/constants"
  "github.com/RrNn/detector/models"
  "github.com/dgrijalva/jwt-go"
  "golang.org/x/crypto/bcrypt"
)

// https://pkg.go.dev/github.com/dgrijalva/jwt-go#DecodeSegment
func CreateJwtToken(user *models.User) (t string, err error) {
  token := jwt.New(jwt.SigningMethodHS256)
  claims := token.Claims.(jwt.MapClaims)
  claims["user_id"] = uint(user.ID)
  claims["name"] = user.Name
  claims["email"] = user.Email
  claims["address"] = user.Address
  claims["contact"] = user.Contact
  claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
  if t, err = token.SignedString([]byte(constants.JwtString)); err != nil {
    return t, err
  }
  return t, nil
}

func GenerateHashFromPassword(password string) (hash string, err error) {
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  if err != nil {
    return "", err
  }
  return string(hashedPassword), nil
}

func CompareHashAndPassword(passwordHash string, password string) (matched bool) {
  if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
    return false
  }
  return true
}

type JwtChecker struct {
}
