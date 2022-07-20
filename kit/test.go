package kit

import (
	"bytes"
	"github.com/dgrijalva/jwt-go"
	"github.com/fede/golang_api/internal/domain/entity"
	"github.com/fede/golang_api/internal/platform/service"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
)

func GetTestConfig(httpMethod string, url string, requestBody []byte, ctxParams []gin.Param) (*httptest.ResponseRecorder, *gin.Context) {

	requestReader := bytes.NewReader(requestBody)
	response := getTargetResponse()
	contextTest := getMockedContext(httpMethod, url, requestReader, response)
	contextTest.Params = ctxParams
	contextTest.Request.Header.Add("Content-Type", "application/json")
	setTokenContext(contextTest)
	return response, contextTest
}

func getMockedContext(method, url string, requestBody io.Reader, response *httptest.ResponseRecorder) *gin.Context {
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(method, url, requestBody)
	return c
}

func getTargetResponse() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}

func setTokenContext(contextTest *gin.Context) {
	serviceJWT := service.NewJWTService()
	r := entity.Role{
		ID:    1,
		Name:  "admin",
		Users: nil,
	}

	u := entity.User{
		ID:       22,
		Name:     "paco",
		Email:    "educabot@gmail.com",
		Password: "cuca",
		Token:    "",
		RoleID:   33,
		Roles:    &r,
	}
	user := serviceJWT.GenerateToken(u)

	token, _ := serviceJWT.ValidateToken(user)
	claims := token.Claims.(jwt.MapClaims)
	contextTest.Set("Claim", claims)
}
