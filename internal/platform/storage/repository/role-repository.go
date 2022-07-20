package repository

import (
	"github.com/fede/golang_api/internal/domain/entity"
	"gorm.io/gorm"
)

type RoleRepository interface {
	InsertRole(user entity.Role) entity.Role
	IsDuplicateRole(role string) (tx *gorm.DB)
	FindByRole(name string) *entity.Role
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

func (db *RoleConnection) InsertRole(role entity.Role) entity.Role {
	db.connection.Save(&role)
	return role
}

func (db *RoleConnection) IsDuplicateRole(name string) (tx *gorm.DB) {
	var role entity.Role
	return db.connection.Where("name = ?", name).Take(&role)
}

func (db *RoleConnection) FindByRole(name string) *entity.Role {
	var role entity.Role
	db.connection.Where("name = ?", name).Take(&role)
	return &role
}
