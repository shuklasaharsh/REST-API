package users_controllers

import (
	"github.com/gofiber/fiber/v2"
	users_service "github.com/shuklasaharsh/REST-API/pkg/REST/services/users"
	users_entity "github.com/shuklasaharsh/REST-API/pkg/entity/users"
	"log"
)

func CreateUser(ctx *fiber.Ctx) error {
	var putUser users_entity.User
	err:= ctx.BodyParser(&putUser)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(&fiber.Map{
			"status": "failed",
			"message": err,
		})
	}
	user, err := users_service.CreateUser(putUser)
	if err != nil {
		log.Println(err)
		return ctx.Status(400).JSON(&fiber.Map{
			"status": "failed",
			"message": err,
		})

	}
	return ctx.Status(201).JSON(&fiber.Map{
		"status": "success",
		"message": "Added user",
		"data": user,
	})
}

func GetUser(ctx *fiber.Ctx) error {
	var id uint
	err := ctx.BodyParser(id)
	if err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"status": "failed",
			"message": err,
		})

	}
	findUser, err := users_service.GetUser(id)
	if err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"status": "failed",
			"message": err,
		})

	}
	return ctx.Status(200).JSON(&fiber.Map{
		"status": "success",
		"message": "findOne",
		"data": findUser,
	})
}

func GetAllUsers(ctx *fiber.Ctx) error {
	users, err := users_service.GetAllUsers()
	if err != nil {
		return fiber.NewError(500, "Service not available")

	}
	return ctx.Status(200).JSON(&fiber.Map{
		"status": "success",
		"message": "findAll",
		"data": users,
	})
}

func UpdateUserByID(ctx *fiber.Ctx) error {
	var updateUser users_entity.User
	err:= ctx.BodyParser(&updateUser)
	if err != nil {
		return fiber.NewError(400, "Service not available")
	}
	user, err := users_service.UpdateUser(updateUser.ID, updateUser.Name, updateUser.Email)
	if err != nil {
		return ctx.Status(404).JSON(&fiber.Map{
			"status": "failed",
			"message": "Not Found",
		})

	}
	return ctx.Status(201).JSON(&fiber.Map{
		"status": "success",
		"message": "updateById",
		"data": user,
	})

}

func DeleteUser(ctx *fiber.Ctx) error {
	var id uint
	err:= ctx.BodyParser(id)
	if err != nil {
		return fiber.NewError(400, "Request Must contain ID")

	}
	err=users_service.DeleteUser(id)
	if err!=nil {
		return fiber.NewError(404, "User not found")

	}
	return ctx.Status(200).JSON(&fiber.Map{
		"status": "success",
		"message": "deleteOne",
	})
}