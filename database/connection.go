package database

import (
	"fmt"

	config "github.com/eufelipemateus/bolierplate-go/configs"
	"github.com/eufelipemateus/bolierplate-go/database/query"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenConnection() {
	conf := config.GetDB()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable ",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Erro while try connect to database.")
	}

	DB = database

	query.SetDefault(DB)
}
