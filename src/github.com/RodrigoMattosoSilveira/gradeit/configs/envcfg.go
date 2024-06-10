package configs

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)


func SetEnv() {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		panic("SetEnv: GOPATH not defined")
	}
	appPath := fmt.Sprintf("%s/src/github.com/RodrigoMattosoSilveira/gradeit", goPath)
	// If ./configs/.env exists, load ./configs/.env
	err := godotenv.Load(fmt.Sprintf("%s/configs/.env", appPath))
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
			appEnvFileName = fmt.Sprintf("%s/configs/.dev.env", appPath)
		case "QA":
			appEnvFileName = fmt.Sprintf("%s/configs/.qa.env", appPath)
		case "E2E":
			appEnvFileName = fmt.Sprintf("%s/configs/.e2e.env", appPath)
		case "STAGE":
			appEnvFileName = fmt.Sprintf("%s/configs/.stage.env", appPath)
		case "PROD":
			appEnvFileName = fmt.Sprintf("%s/configs/.prod.env", appPath)
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
	os.Setenv("DB_NAME", fmt.Sprintf("%s/%s", appPath, os.Getenv("DB_NAME")))
	slog.Info(fmt.Sprintf("DB_NAME = %s",  os.Getenv("DB_NAME")))
}