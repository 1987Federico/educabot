package repository

import (
	"github.com/fede/golang_api/internal/domain/entity"
	"github.com/fede/golang_api/internal/platform/utils"
	"gorm.io/gorm"
)

//UserRepository is contract what userRepository can do to db
type UserRepository interface {
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.User
	ProfileUser(userID string) entity.User
}

type UserConnection struct {
	connection *gorm.DB
}

//NewUserRepository is creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) *UserConnection {
	return &UserConnection{
		connection: db,
	}
}

func (db *UserConnection) InsertUser(user entity.User) entity.User {
	user.Password = utils.HashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	db.connection.Preload("Roles").Find(&user)
	return user
}

func (db *UserConnection) UpdateUser(user entity.User) entity.User {
	if user.Password != "" {
		user.Password = utils.HashAndSalt([]byte(user.Password))
	} else {
		var tempUser entity.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	db.connection.Save(&user)
	return user
}

func (db *UserConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := db.connection.Preload("Roles").Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *UserConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Preload("Roles").Where("email = ?", email).Take(&user)
}

func (db *UserConnection) FindByEmail(email string) entity.User {
	var user entity.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *UserConnection) ProfileUser(userID string) entity.User {
	var user entity.User
	db.connection.Preload("Roles").Find(&user, userID)
	return user
}
