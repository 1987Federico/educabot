package service

import (
	"fmt"
	"github.com/fede/golang_api/internal/domain/entity"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//JWTService is a contract of what jwtService can do
type JWTService interface {
	GenerateToken(user entity.User) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

type JwtServices struct {
	secretKey string
	issuer    string
}

//NewJWTService method is creates a new instance of JWTService
func NewJWTService() *JwtServices {
	return &JwtServices{
		issuer:    "ydhnwb",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "ydhnwb"
	}
	return secretKey
}

func (j *JwtServices) GenerateToken(user entity.User) string {
	claims := &jwtCustomClaim{
		fmt.Sprintf("%d", user.ID),
		user.Roles.Name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *JwtServices) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
