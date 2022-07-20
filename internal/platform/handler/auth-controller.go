package handler

import (
	"github.com/dgrijalva/jwt-go"
	dto2 "github.com/fede/golang_api/internal/domain/dto"
	"github.com/fede/golang_api/internal/domain/entity"
	"github.com/fede/golang_api/internal/platform/helper/errors"
	"github.com/fede/golang_api/internal/platform/helper/response"
	service2 "github.com/fede/golang_api/internal/platform/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//AuthController interface is a contract what this handler can do
type AuthController interface {
	Login(ctx *gin.Context)
	RegisterDrive(ctx *gin.Context)
}

type AuthControllers struct {
	authService *service2.AuthServices
	jwtService  *service2.JwtServices
}

//NewAuthController creates a new instance of AuthController
func NewAuthController(authService *service2.AuthServices, jwtService *service2.JwtServices) *AuthControllers {
	return &AuthControllers{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c AuthControllers) Login(ctx *gin.Context) {
	var loginDTO dto2.LoginDTO

	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		panic(errors.BadRequestApiError("Failed to process request", errDTO.Error()))
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	v, ok := authResult.(entity.User)
	if !ok {
		panic(errors.StatusUnauthorizedApiError("Please check again your credential", "Invalid Credential"))
	}
	generatedToken := c.jwtService.GenerateToken(v)
	v.Token = generatedToken
	response := response.BuildResponse(true, "OK!", v)
	ctx.JSON(http.StatusOK, response)
	return

}

func (c AuthControllers) RegisterDrive(ctx *gin.Context) {
	var registerDTO dto2.RegisterDTO
	token, _ := ctx.Get("Claim")
	claims := token.(jwt.MapClaims)
	if claims["role"] != "admin" {
		panic(errors.ForbiddenApiError("Failed to process request", "User not authorized to perform this action"))
	}
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		panic(errors.BadRequestApiError("Failed to process request", errDTO.Error()))
	}

	role := c.authService.RoleExist(registerDTO.Role)
	c.authService.IsDuplicateEmail(registerDTO.Email)
	c.authService.DriverExist(registerDTO.Driver.DriverFile)

	createdUser := c.authService.CreateUser(registerDTO, role.ID)
	tokenUser := c.jwtService.GenerateToken(createdUser)
	createdUser.Token = tokenUser
	response := response.BuildResponse(true, "OK!", createdUser)
	ctx.JSON(http.StatusCreated, response)
}
