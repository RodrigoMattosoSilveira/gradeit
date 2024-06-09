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
	if appEnv == "" {
		appEnv = "DEV"
	}
	var appEnvFileName string
	switch appEnv {
		case "DEV":
			appEnvFileName = "./configs/.dev.env"
		case "QA":
			appEnvFileName = "./configs/.qa.env"
		case "E2E":
			appEnvFileName = "./configs/.e2e.env"
		case "STAGE":
			appEnvFileName = "./configs/.stage.env"
		case "PROD":
			appEnvFileName = "./configs/.prod.env"
	}

	err = godotenv.Overload(appEnvFileName)
	if err != nil {
		panic(fmt.Sprintf("SetEnv: error loading%s", appEnvFileName))
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