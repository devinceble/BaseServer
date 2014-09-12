package TestsConfigs

import (
	"testing"

	. "github.com/devinceble/BaseServer/Configs"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Configuration(t *testing.T) {
	Convey("Connection String Test", t, func() {
		var conf Config
		So(
			conf.Connection(),
			ShouldResemble,
			"root:qwerty@unix(/var/run/mysqld/mysqld.sock)/test_bss?charset=utf8",
		)
		So(conf.Dbase(), ShouldResemble, "mysql")
	})
}
