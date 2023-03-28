package db

import (
	"fmt"
	"os"

	"github.com/mjawa20/todo-list-go.git/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type db struct {
	db *gorm.DB
}

func NewPostgres() domain.DB {
	db := db{}
	db.initConnection()
	return &db
}

func (d *db) initConnection() {
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)

	var err error
	d.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	d.db = d.db.Debug()
	if err != nil {
		panic(err)
	}

	d.migrate()
}

func (d *db) GetConnection() *gorm.DB {
	if d.db == nil {
		d.initConnection()
	}
	return d.db
}

func (d *db) migrate() {
	d.db.AutoMigrate(
		domain.Activity{},
		domain.Todo{},
	)
}
