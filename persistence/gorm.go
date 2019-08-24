package persistence

import (
	"github.com/jinzhu/gorm"
)

type IGormDB interface {
	Find(out interface{}, where ...interface{}) IGormDB
	Save(value interface{}) IGormDB
}

type gormWrapper struct {
	db *gorm.DB
}

func (gw *gormWrapper) Find(out interface{}, where ...interface{}) IGormDB {
	gw.db = gw.db.Find(out, where...)
	return gw
}

func (gw *gormWrapper) Save(value interface{}) IGormDB {
	gw.db = gw.db.Save(value)
	return gw
}

func Wrap(db *gorm.DB) *gormWrapper {
	return &gormWrapper{db: db}
}
