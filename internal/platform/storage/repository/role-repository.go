package repository

import (
	"context"
	"github.com/fede/golang_api/internal/domain/entity"
	"github.com/fede/golang_api/internal/platform/storage"
	"gorm.io/gorm"
)

type RoleRepository interface {
	InsertRole(user entity.Role, ctx context.Context) entity.Role
	IsDuplicateRole(role string, ctx context.Context) error
	FindByRole(name string, ctx context.Context) *entity.Role
}

type RoleConnection struct {
	connection *gorm.DB
}

//NewRoleRepository is creates a new instance of RoleRepository
func NewRoleRepository(db *gorm.DB) *RoleConnection {
	return &RoleConnection{
		connection: db,
	}
}

func (r *RoleConnection) InsertRole(role entity.Role, ctx context.Context) entity.Role {
	db := storage.FromContext(ctx)
	db.Save(&role)
	return role
}

func (r *RoleConnection) IsDuplicateRole(name string, ctx context.Context) error {
	var role entity.Role
	db := storage.FromContext(ctx)
	err := db.Where("name = ?", name).Take(&role).Error
	return err
}

func (r *RoleConnection) FindByRole(name string, ctx context.Context) *entity.Role {
	var role entity.Role
	db := storage.FromContext(ctx)
	db.Where("name = ?", name).Take(&role)
	return &role
}
