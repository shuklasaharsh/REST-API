package users_entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID uint `gorm:"default:uuid_generate_v4();primaryKey" json:"id"`
	Name string
	Email string
}