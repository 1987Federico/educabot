package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/fede/golang_api/internal/platform/helper/errorCustom"
	"github.com/fede/golang_api/internal/platform/service"
	"github.com/gin-gonic/gin"
	"log"
)

type AuthJWT struct {
	jwtService *service.JwtServices
}

func NewAuthorizeJWT(jwtService *service.JwtServices) *AuthJWT {
	return &AuthJWT{
		jwtService: jwtService,
	}
}

//AuthorizeJWT validates the token user given, return 401 if not valid
func (auth *AuthJWT) AuthorizeJWT(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		panic(errorCustom.BadRequestApiError("Failed to process request", "No token found"))
		return
	}
	token, err := auth.jwtService.ValidateToken(authHeader)
	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		c.Set("Claim", claims)
		log.Println("Claim[user_id]: ", claims["user_id"])
		log.Println("Claim[issuer] :", claims["issuer"])
	} else {
		log.Println(err)
		panic(errorCustom.StatusUnauthorizedApiError("Token is not valid", err.Error()))
	}
}
