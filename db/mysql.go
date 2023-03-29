package db

import (
	"fmt"
	"os"

	"github.com/mjawa20/todo-list-go.git/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type db struct {
	db *gorm.DB
}

func NewMysql() domain.DB {
	db := db{}
	db.initConnection()
	return &db
}

func (d *db) initConnection() {
	dbName := os.Getenv("MYSQL_DBNAME")
	dbPort := os.Getenv("MYSQL_PORT")
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println(dsn)
	var err error
	d.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
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
		domain.Activities{},
		domain.Todos{},
	)
}
