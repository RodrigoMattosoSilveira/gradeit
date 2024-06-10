package configs

import (
	"gorm.io/gorm"
  )

var DB *gorm.DB

func Config () {
	SetEnv()
	DBInit()
	DoAutomigrate()
	
}