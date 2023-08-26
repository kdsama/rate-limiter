package utils

import (
	"errors"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyz1234567890-")
)

var (
	Err_InvalidToken            = errors.New("invalid token")
	Err_UserIdMissing           = errors.New("userid is missing in the token")
	Err_UnexpectedSigningMethod = errors.New("unexpected signing method")
)

// jwt code has been copied from chatgpt
// Define a secret key used to sign and verify the JWT tokens
var secretKey = []byte("my-secret-key")

// Create a JWT token with a given user ID and expiration time
func CreateJWTToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Verify a JWT token and extract the user ID
func VerifyJWTToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, Err_UnexpectedSigningMethod
		}

		return secretKey, nil
	})

	if err != nil {
		return "", Err_InvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", Err_InvalidToken
	}

	userID, ok := claims["userid"].(string)
	if !ok {
		return "", Err_UserIdMissing
	}

	return userID, nil
}

func GenerateHashForPassword(input string) (string, error) {
	password := []byte(input)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(input string, password string) error {

	err := bcrypt.CompareHashAndPassword([]byte(input), []byte(password))
	return err
}

func GenerateUUID() string {
	return uuid.New().String()

}

func GenerateShortURL() string {
	rand.Seed(time.Now().UnixMicro())
	return randSeq(rand.Intn(6))
}
func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
