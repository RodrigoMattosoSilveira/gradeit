package configs

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)


func SetEnv() error {
	// If ./configs/.env exists, load ./configs/.env
	err := godotenv.Load("./configs/.env")
	if err != nil {
		log.Panicln("error loading .env file")
		return err
	}

	appEnv := os.Getenv("APP_ENV")

	if appEnv=="DEV" {
		err = godotenv.Overload("./configs/.dev.env")
		if err != nil {
			log.Printf(("error loading .dev.env file"))
			return err
		}
	} else {
		if appEnv=="STAGE" {
			err = godotenv.Overload("./configs/.stage.env")
			if err != nil {
				log.Printf(("error loading .stage.env file"))
				return err
			}
		} else {
			if appEnv=="PROD" {
				err = godotenv.Overload("./configs/.prod.env")
				if err != nil {
					log.Printf(("error loading .prod.env file"))
					return err
				}
			}
		}
	}
	return err
}