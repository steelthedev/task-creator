package conn

import (
	"os"

	"github.com/steelthedev/task-handler/data"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	url := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              false,
	})

	if err != nil {
		return err
	}

	DB.AutoMigrate(data.Task{})

	return nil
}
