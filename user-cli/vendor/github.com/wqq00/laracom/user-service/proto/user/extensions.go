package laracom_user_service

import (
	"github.com/jinzhu/gorm"
	"github.com/google/uuid"
)


func (model *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.New()
	return scope.SetColumn("Id", uuid.String())
}