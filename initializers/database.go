package initializers

import (
	"os"

	"github.com/captainmio/go-crud-post/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {

	var err error
	dsn := os.Getenv("DSN")

	db.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}
}

// psql 'postgresql://neondb_owner:npg_zMayAwcX7ST3@ep-long-tooth-adlyks17-pooler.c-2.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require'
