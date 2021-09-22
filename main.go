package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/viper"
	database "github.com/shuklasaharsh/REST-API/pkg/database"
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
	_, err := database.ConnectDB(url)
	if err!= nil {
		log.Panicln(fmt.Errorf("Database could not be connected: %w ", err))
	}

}
