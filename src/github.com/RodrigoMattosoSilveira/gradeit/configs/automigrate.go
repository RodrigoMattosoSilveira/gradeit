package configs

import "github.com/RodrigoMattosoSilveira/gradeit/models"

func DoAutomigrate () {
	DB.AutoMigrate(&models.Person{})
	DB.AutoMigrate(&models.Assignment{})
}