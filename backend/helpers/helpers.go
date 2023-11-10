package helpers

import (
	"backend/config"
	"backend/models"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthCustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func PasswordHashing(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}
	hash := string(hashedPassword)
	return hash, nil
}

func GenerateTokenUsers(userID int, userEmail string, expirationTime time.Time) (string, error) {
	cfg, _ := config.LoadConfig()

	claims := &AuthCustomClaims{
		Id:    userID,
		Email: userEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.KEY))

	// tokenString, err:=jwt.NewWithClaims(jwt.SigningMethodHS256, &AuthCustomClaims{Id:    userID, Email: userEmail,StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix(),IssuedAt:  time.Now().Unix(),},}).SignedString([]byte(cfg.KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateAccessToken(user models.SignupDetailResponse) (string, error) {

	expirationTime := time.Now().Add(15 * time.Minute)
	tokenString, err := GenerateTokenUsers(user.Id, user.Email, expirationTime)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func GenerateRefreshToken(user models.SignupDetailResponse) (string, error) {

	expirationTime := time.Now().Add(24 * 90 * time.Hour)
	tokeString, err := GenerateTokenUsers(user.Id, user.Email, expirationTime)
	if err != nil {
		return "", err
	}
	return tokeString, nil

}

func GetTokenFromHeader(header string) string {
	// Example header format: "Bearer <token>"

	if len(header) > 7 && header[:7] == "Bearer " {
		return header[7:]
	}

	return header
	// return ""
}

func ExtractUserIDFromToken(tokenString string) (int, string, error) {
	cfg, _ := config.LoadConfig()

	token, err := jwt.ParseWithClaims(tokenString, &AuthCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte(cfg.KEY), nil // Replace with your actual secret key
	})

	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*AuthCustomClaims)
	if !ok {
		return 0, "", fmt.Errorf("invalid token claims")
	}

	return claims.Id, claims.Email, nil

}
