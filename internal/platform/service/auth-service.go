package service

import (
	"context"
	"errors"
	"github.com/fede/golang_api/internal/domain/dto"
	entity2 "github.com/fede/golang_api/internal/domain/entity"
	"github.com/fede/golang_api/internal/platform/helper/errorCustom"
	repository2 "github.com/fede/golang_api/internal/platform/storage/repository"
	"gorm.io/gorm"
	"log"

	"golang.org/x/crypto/bcrypt"
)

//AuthService is a contract about something that this service can do
type AuthService interface {
	VerifyCredential(email string, password string, ctx context.Context) interface{}
	CreateUser(user dto.RegisterDTO, roleID uint64, ctx context.Context) entity2.User
	FindByEmail(email string, ctx context.Context) entity2.User
	IsDuplicateEmail(email string, ctx context.Context) bool
	RoleExist(name string, ctx context.Context) *entity2.Role
	DriverExist(driverFile uint64, ctx context.Context) bool
}

type AuthServices struct {
	userRepository   *repository2.UserConnection
	roleRepository   *repository2.RoleConnection
	driverRepository *repository2.DriverConnection
}

//NewAuthService creates a new instance of AuthService
func NewAuthService(userRep *repository2.UserConnection, roleRep *repository2.RoleConnection, driverRep *repository2.DriverConnection) *AuthServices {
	return &AuthServices{
		userRepository:   userRep,
		roleRepository:   roleRep,
		driverRepository: driverRep,
	}
}

func (service *AuthServices) VerifyCredential(email string, password string, ctx context.Context) interface{} {
	res := service.userRepository.VerifyCredential(email, password, ctx)
	if v, ok := res.(entity2.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *AuthServices) CreateUser(user dto.RegisterDTO, roleID uint64, ctx context.Context) entity2.User {
	userToCreate := entity2.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Token:    "",
		RoleID:   roleID,
		Driver: &entity2.Driver{
			DriverFile:  user.Driver.DriverFile,
			Description: user.Driver.Description,
		},
	}
	res := service.userRepository.InsertUser(userToCreate, ctx)
	return res
}

func (service *AuthServices) FindByEmail(email string, ctx context.Context) entity2.User {
	return service.userRepository.FindByEmail(email, ctx)
}

func (service *AuthServices) IsDuplicateEmail(email string, ctx context.Context) bool {
	err := service.userRepository.IsDuplicateEmail(email, ctx)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(errorCustom.ConflictApiError("Failed to process request", "Duplicate email"))
	}
	return false
}

func (service *AuthServices) RoleExist(name string, ctx context.Context) *entity2.Role {
	res := service.roleRepository.FindByRole(name, ctx)
	if res == nil {
		panic(errorCustom.ConflictApiError("Failed to process request", "Role not Exist"))
	}
	return res
}

func (service *AuthServices) DriverExist(driverFile uint64, ctx context.Context) bool {
	res := service.driverRepository.FindDriverByFile(driverFile, ctx)
	if res != nil {
		panic(errorCustom.ConflictApiError("Failed to process request", "Driver already Exist"))
	}
	return false
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
