package TestsModels

import (
	"testing"

	// . "github.com/devinceble/BaseServer/Configs"
	. "github.com/devinceble/BaseServer/Models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Database(t *testing.T) {
	Convey("Database Connetion Test", t, func() {
		var db gorm.DB
		db, err := Connect()
		So(err, ShouldBeNil)
		err2 := db.DB().Ping()
		So(err2, ShouldBeNil)
	})
}
