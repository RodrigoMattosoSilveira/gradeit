package configs

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)


func SetEnv() {
	// If ./configs/.env exists, load ./configs/.env
	err := godotenv.Load("./configs/.env")
	if err != nil {
		panic(fmt.Sprintf("SetEnv: error loading %s file", ".env"))
	}

	appEnv := os.Getenv("APP_ENV")

	if appEnv=="DEV" {
		err = godotenv.Overload("./configs/.dev.env")
		if err != nil {
			panic("SetEnv: error loading .dev.env file")
		}
	} else {
		if appEnv=="STAGE" {
			err = godotenv.Overload("./configs/.stage.env")
			if err != nil {
				panic("SetEnv: error loading .stage.env file")
			}
		} else {
			if appEnv=="PROD" {
				err = godotenv.Overload("./configs/.prod.env")
				if err != nil {
					panic("SetEnv: error loading .prod.env file")
				}
			}
		}
	}

	// Showcase environment variables handling
	// without setting the environment
	//
	slog.Info("SetEnv: environment setup successfully")
	slog.Info(fmt.Sprintf("THIS_ENV = %s",  os.Getenv("THIS_ENV")))
	slog.Info(fmt.Sprintf("HTTP_PORT = %s",  os.Getenv("HTTP_PORT")))
	slog.Info(fmt.Sprintf("DB_DIALECT = %s",  os.Getenv("DB_DIALECT")))
	slog.Info(fmt.Sprintf("DB_NAME = %s",  os.Getenv("DB_NAME")))

}