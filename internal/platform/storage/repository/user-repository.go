package repository

import (
	"context"
	"github.com/fede/golang_api/internal/domain/entity"
	"github.com/fede/golang_api/internal/platform/storage"
	"github.com/fede/golang_api/internal/platform/utils"
	"gorm.io/gorm"
)

//UserRepository is contract what userRepository can do to db
type UserRepository interface {
	InsertUser(user entity.User, ctx context.Context) entity.User
	UpdateUser(user entity.User, ctx context.Context) entity.User
	VerifyCredential(email string, password string, ctx context.Context) interface{}
	IsDuplicateEmail(email string, ctx context.Context) error
	FindByEmail(email string, ctx context.Context) entity.User
	ProfileUser(userID string, ctx context.Context) entity.User
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

func (u *UserConnection) InsertUser(user entity.User, ctx context.Context) entity.User {
	user.Password = utils.HashAndSalt([]byte(user.Password))
	db := storage.FromContext(ctx)
	db.Save(&user)
	db.Preload("Roles").Find(&user)
	return user
}

func (u *UserConnection) UpdateUser(user entity.User, ctx context.Context) entity.User {
	db := storage.FromContext(ctx)
	if user.Password != "" {
		user.Password = utils.HashAndSalt([]byte(user.Password))
	} else {
		var tempUser entity.User
		db.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	db.Save(&user)
	return user
}

func (u *UserConnection) VerifyCredential(email string, password string, ctx context.Context) interface{} {
	var user entity.User
	db := storage.FromContext(ctx)
	res := db.Preload("Roles").Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (u *UserConnection) IsDuplicateEmail(email string, ctx context.Context) error {
	var user entity.User
	db := storage.FromContext(ctx)
	err := db.Preload("Roles").Where("email = ?", email).Take(&user).Error
	return err
}

func (u *UserConnection) FindByEmail(email string, ctx context.Context) entity.User {
	var user entity.User
	db := storage.FromContext(ctx)
	db.Where("email = ?", email).Take(&user)
	return user
}

func (u *UserConnection) ProfileUser(userID string, ctx context.Context) entity.User {
	var user entity.User
	db := storage.FromContext(ctx)
	db.Preload("Roles").Find(&user, userID)
	return user
}
