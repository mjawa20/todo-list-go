package domain

import "gorm.io/gorm"

type DB interface {
	GetConnection() *gorm.DB
}
