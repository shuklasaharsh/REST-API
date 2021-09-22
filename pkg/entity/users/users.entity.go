package users_entity

type User struct {
	ID int `gorm:"default:uuid_generate_v4();primaryKey;type:bigint" json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}