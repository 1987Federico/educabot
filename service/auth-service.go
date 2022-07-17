package service

import (
	"log"

	"github.com/ydhnwb/golang_api/dto"
	"github.com/ydhnwb/golang_api/entity"
	"github.com/ydhnwb/golang_api/repository"
	"golang.org/x/crypto/bcrypt"
)

//AuthService is a contract about something that this service can do
type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO, roleID uint64) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
	RoleExist(name string) *entity.Role
	DriverExist(driverFile uint64) bool
}

type authService struct {
	userRepository   repository.UserRepository
	roleRepository   repository.RoleRepository
	driverRepository repository.DriverRepository
}

//NewAuthService creates a new instance of AuthService
func NewAuthService(userRep repository.UserRepository, roleRep repository.RoleRepository, driverRep repository.DriverRepository) AuthService {
	return &authService{
		userRepository:   userRep,
		roleRepository:   roleRep,
		driverRepository: driverRep,
	}
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user dto.RegisterDTO, roleID uint64) entity.User {
	userToCreate := entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Token:    "",
		RoleID:   roleID,
		Driver: &entity.Driver{
			DriverFile:  user.Driver.DriverFile,
			Description: "",
		},
	}
	//err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	//userToCreate.RoleID = roleID
	//if err != nil {
	//	log.Fatalf("Failed map %v", err)
	//}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}

func (service *authService) FindByEmail(email string) entity.User {
	return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func (service *authService) RoleExist(name string) *entity.Role {
	res := service.roleRepository.FindByRole(name)
	return res
}

func (service *authService) DriverExist(driverFile uint64) bool {
	res := service.driverRepository.FindDriverByFile(driverFile)
	return res == nil
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
