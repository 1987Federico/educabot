package service

import (
	"github.com/fede/golang_api/internal/domain/dto"
	"github.com/fede/golang_api/internal/domain/entity"
	"github.com/fede/golang_api/internal/platform/storage/repository"
	"log"

	"github.com/mashingan/smapping"
)

//UserService is a contract.....
type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
}

type UserServices struct {
	userRepository *repository.UserConnection
}

//NewUserService creates a new instance of UserService
func NewUserService(userRepo *repository.UserConnection) *UserServices {
	return &UserServices{
		userRepository: userRepo,
	}
}

func (service *UserServices) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *UserServices) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}
