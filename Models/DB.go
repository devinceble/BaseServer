package Models

import (
	"github.com/devinceble/BaseServer/Configs"
	"github.com/jinzhu/gorm"
)

//DB Structure
type DB struct {
	db gorm.DB
}

//Mdb DB Variable
var Mdb DB

//Connect Method
func Connect() (gorm.DB, error) {
	var conf Configs.Config
	db, err := gorm.Open(conf.Dbase(), conf.Connection())
	if err != nil {
		return db, err
	}
	db.LogMode(false)
	db.DB()

	// Then you could invoke `*sql.DB`'s functions with it
	db.DB().Ping()
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	// Disable table name's pluralization
	db.SingularTable(true)
	return db, nil
}
