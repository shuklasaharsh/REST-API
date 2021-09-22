package users_service

import (
	"errors"
	"fmt"
	"github.com/shuklasaharsh/REST-API/pkg/database"
	"github.com/shuklasaharsh/REST-API/pkg/entity/users"
)

func CreateUser(putUser users_entity.User) (*users_entity.User, error){
	db:= database.GetDB()
	fmt.Println(putUser)
	err:= db.Create(&putUser).Error
	if err != nil {
		fmt.Errorf("\nerr: %w ", err)
		return nil, err
	}

	return &putUser, nil

}

func GetUser(id uint) (*users_entity.User, error) {
	db:= database.GetDB()
	var user users_entity.User

	err:= db.Model(users_entity.User{}).Where("id=?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAllUsers() ([]*users_entity.User, error) {
	db := database.GetDB()
	var users[] *users_entity.User
	err := db.Model(users_entity.User{}).Find(&users).Error

	if err != nil {
		return nil, err
	}
	return users, nil
}

func UpdateUser(id int, name string, email string) (*users_entity.User, error) {
	db:= database.GetDB()
	var user users_entity.User
	err := db.Model(users_entity.User{}).Where("id=?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, errors.New("404")
	}
	user.Name = name
	user.Email = email

	err = db.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUser(id uint) error {
	db:= database.GetDB()
	err := db.Delete("id=?", id).Delete(&users_entity.User{}).Error
	if err!= nil {
		return err
	}
	return nil
}