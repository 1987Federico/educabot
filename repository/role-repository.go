package repository

import (
	"github.com/ydhnwb/golang_api/entity"
	"gorm.io/gorm"
)

type RoleRepository interface {
	InsertRole(user entity.Role) entity.Role
	IsDuplicateRole(role string) (tx *gorm.DB)
	FindByRole(name string) *entity.Role
}

type roleConnection struct {
	connection *gorm.DB
}

//NewRoleRepository is creates a new instance of RoleRepository
func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleConnection{
		connection: db,
	}
}

func (db *roleConnection) InsertRole(role entity.Role) entity.Role {
	db.connection.Save(&role)
	return role
}

func (db *roleConnection) IsDuplicateRole(name string) (tx *gorm.DB) {
	var role entity.Role
	return db.connection.Where("name = ?", name).Take(&role)
}

func (db *roleConnection) FindByRole(name string) *entity.Role {
	var role entity.Role
	db.connection.Where("name = ?", name).Take(&role)
	return &role
}
