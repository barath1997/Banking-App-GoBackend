package helpers

import (
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"google.com/banking-app/models"
)

type Errors interface {
	HandleErr()
}

type Error struct {
	ErrorMessage string
	Err          error
}

func (e *Error) HandleErr() {

	log.Printf("%s : %s\n", e.ErrorMessage, e.Err.Error())

}

func HashandSalt(pass []byte) string {

	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		er := &Error{"unable to hash password", err}
		er.HandleErr()
	}

	return string(hashed)
}

func GetJwtToken(userId string) string {

	jwtClaims := jwt.MapClaims{
		"user_id": userId,
		"expiry":  time.Now().Add(time.Minute * 10).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	token, err := jwtToken.SignedString([]byte("secret signature"))
	if err != nil {
		er := &Error{ErrorMessage: "cannot sign jwt token", Err: err}
		er.HandleErr()
	}

	return token
}

func Validation(values []models.Validation) (string, bool) {

	username := regexp.MustCompile(`^([A-Za-z0-9. ]{5,})+$`)
	email := regexp.MustCompile(`^[A-Za-z0-9.]+[@]+[A-Za-z0-9]+[.]+[A-Za-z0-9]+$`)

	for _, val := range values {
		switch val.Valid {
		case "username":
			if !username.MatchString(val.Value) {
				return "invalid username", false
			}

		case "email":
			if !email.MatchString(val.Value) {
				return "invalid email", false
			}
		case "password":
			if len(val.Value) < 5 {
				return "invalid password", false
			}
		default:
			log.Panicf("incorrect object : %s\n", val.Valid)
			return "incorrect object", false
		}
	}
	return "valid", true
}

func ValidateToken(id, jwtToken string) bool {

	cleanJwt := strings.Replace(jwtToken, "Bearer", "", -1)
	sp := strings.TrimSpace(cleanJwt)
	tokenData := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(sp, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret signature"), nil
	})
	if err != nil {
		er := &Error{ErrorMessage: "error in validating jwt", Err: err}
		er.HandleErr()
	}

	// validation
	if token.Valid && tokenData["user_id"] == id {
		return true
	} else {
		return false
	}

}
