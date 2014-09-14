package Migrations

import (
	"github.com/devinceble/BaseServer/Helpers"
	. "github.com/devinceble/BaseServer/Models"
)

func Migrate() {
	var user User
	var profile Profile
	var address Address
	var office Office
	var email Email
	var phone Phone
	Helpers.BaseLog("DATABASE", "SCHEMA", "", "PRODUCTION", 0, 0, nil)

	Helpers.BaseLog("DATABASE", "TABLE", "", "USER", 1, 0, nil)
	user.Migrate()

	Helpers.BaseLog("DATABASE", "TABLE", "", "PROFILE", 2, 0, nil)
	profile.Migrate()

	Helpers.BaseLog("DATABASE", "TABLE", "", "ADDRESS", 3, 0, nil)
	address.Migrate()

	Helpers.BaseLog("DATABASE", "TABLE", "", "OFFICE", 4, 0, nil)
	office.Migrate()

	Helpers.BaseLog("DATABASE", "TABLE", "", "EMAIL", 5, 0, nil)
	email.Migrate()

	Helpers.BaseLog("DATABASE", "TABLE", "", "PHONE", 6, 0, nil)
	phone.Migrate()
}
