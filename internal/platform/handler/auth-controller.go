package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/fede/golang_api/internal/domain/dto"
	"github.com/fede/golang_api/internal/domain/entity"
	"github.com/fede/golang_api/internal/platform/helper/errorCustom"
	"github.com/fede/golang_api/internal/platform/helper/response"
	"github.com/fede/golang_api/internal/platform/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//AuthController interface is a contract what this handler can do
type AuthController interface {
	Login(ctx *gin.Context)
	RegisterDrive(ctx *gin.Context)
}

type AuthControllers struct {
	authService *service.AuthServices
	jwtService  *service.JwtServices
}

//NewAuthController creates a new instance of AuthController
func NewAuthController(authService *service.AuthServices, jwtService *service.JwtServices) *AuthControllers {
	return &AuthControllers{
		authService: authService,
		jwtService:  jwtService,
	}
}

// Login godoc
// @Summary login of users
// @Description allows a user to authenticate
// @Tags Drivers
// @Accept  json
// @Produce  json
// @Param paramsToSearch body dto.RegisterDTO true "driver to register"
// @Success 201 {object} dto.LoginDTO
// @Failure 400 {object} errorCustom.ApiError
// @Failure 401 {object} errorCustom.ApiError
// @Failure 500 {object} errorCustom.ApiError
// @Router /api/auth/login [post]
func (c AuthControllers) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO

	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		panic(errorCustom.BadRequestApiError("Failed to process request", errDTO.Error()))
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password, ctx.Request.Context())
	v, ok := authResult.(entity.User)
	if !ok {
		panic(errorCustom.StatusUnauthorizedApiError("Please check again your credential", "Invalid Credential"))
	}
	generatedToken := c.jwtService.GenerateToken(v)
	v.Token = generatedToken
	res := response.BuildResponse(true, "OK!", v)
	ctx.JSON(http.StatusOK, res)
	return

}

// RegisterDrive godoc
// @Summary register a driver
// @Description register a driver and return it
// @Tags Drivers
// @Accept  json
// @Produce  json
// @Param paramsToSearch body dto.RegisterDTO true "driver to register"
// @Success 201 {object} entity.Driver
// @Failure 400 {object} errorCustom.ApiError
// @Failure 401 {object} errorCustom.ApiError
// @Failure 404 {object} errorCustom.ApiError
// @Failure 409 {object} errorCustom.ApiError
// @Failure 500 {object} errorCustom.ApiError
// @Router /api/driver/register/driver [post]
func (c AuthControllers) RegisterDrive(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	token, _ := ctx.Get("Claim")
	claims := token.(jwt.MapClaims)
	if claims["role"] != "admin" {
		panic(errorCustom.ForbiddenApiError("Failed to process request", "User not authorized to perform this action"))
	}
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		panic(errorCustom.BadRequestApiError("Failed to process request", errDTO.Error()))
	}

	role := c.authService.RoleExist(registerDTO.Role, ctx.Request.Context())
	c.authService.IsDuplicateEmail(registerDTO.Email, ctx.Request.Context())
	c.authService.DriverExist(registerDTO.Driver.DriverFile, ctx.Request.Context())

	createdUser := c.authService.CreateUser(registerDTO, role.ID, ctx.Request.Context())
	tokenUser := c.jwtService.GenerateToken(createdUser)
	createdUser.Token = tokenUser
	res := response.BuildResponse(true, "OK!", createdUser)
	ctx.JSON(http.StatusCreated, res)
}
