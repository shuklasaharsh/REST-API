package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
	usersroutes "github.com/shuklasaharsh/REST-API/pkg/REST/routes/users"
	database "github.com/shuklasaharsh/REST-API/pkg/database"
	users_entity "github.com/shuklasaharsh/REST-API/pkg/entity/users"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err:= viper.ReadInConfig(); err != nil {
		r:= color.New(color.BgRed)
		r.Println("Cannot find .env")
	}
	url, ok := viper.Get("DATABASE_URL").(string)
	if !ok {
		log.Panicln(fmt.Errorf("DB URL was not found"))
	}
	db, err := database.ConnectDB(url)
	err= db.AutoMigrate(&users_entity.User{})
	if err != nil {
		log.Panicln(fmt.Errorf("Error Migrating Models: %w ", err))
	}
	if err!= nil {
		log.Panicln(fmt.Errorf("Database could not be connected: %w ", err))
	}
	app:=fiber.New()
	usersroutes.InitialiseRoutes(app)
	log.Fatalln(app.Listen(":3000"))

}
