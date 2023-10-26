package drivers

import (
	"fmt"
	"code-ai/models/domain"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)


type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func LoadConfig() Config {

	return Config{
		DB_Username: "root",
		DB_Password: "",
		DB_Port:     "3306",
		DB_Host:     "127.0.0.1",
		DB_Name:     "mini_project",
	}

}
  
  
var DB *gorm.DB

func Init() {
	InitDB()
	InitialMigration()
}

func InitDB() {

	config := LoadConfig()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&domain.Message{}, &domain.User{})
}