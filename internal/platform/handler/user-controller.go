package handler

import (
	"fmt"
	"github.com/fede/golang_api/internal/domain/dto"
	"github.com/fede/golang_api/internal/platform/helper/errors"
	"github.com/fede/golang_api/internal/platform/helper/response"
	"github.com/fede/golang_api/internal/platform/service"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//UserController is a ....
type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type UserControllers struct {
	userService *service.UserServices
}

//NewUserController is creating anew instance of UserControlller
func NewUserController(userService *service.UserServices) *UserControllers {
	return &UserControllers{
		userService: userService,
	}
}

func (c *UserControllers) Update(ctx *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := ctx.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		panic(errors.BadRequestApiError("Failed to process request", errDTO.Error()))
	}

	token, _ := ctx.Get("Claim")
	claims := token.(jwt.MapClaims)
	if claims["role"] != "admin" && claims["role"] != "driver" {
		panic(errors.ForbiddenApiError("Failed to process request", "User not authorized to perform this action"))
	}

	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = id
	u := c.userService.Update(userUpdateDTO)
	res := response.BuildResponse(true, "OK!", u)
	ctx.JSON(http.StatusOK, res)
}

func (c *UserControllers) Profile(context *gin.Context) {
	token, _ := context.Get("Claim")
	claims := token.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	if claims["role"] != "admin" {
		panic(errors.ForbiddenApiError("Failed to process request", "User not authorized to perform this action"))
	}
	user := c.userService.Profile(id)
	res := response.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)

}
