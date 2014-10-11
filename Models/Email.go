package Models

import (
	"time"

	"github.com/devinceble/BaseServer/Helpers"
)

//Email Model
type Email struct {
	Id        int64
	ProfileId int64
	Type      string `sql:"type:varchar(20);not null"`
	Email     string

	Crat int64
	Upat int64
	Dlat int64
}

//Find Email
func (email *Email) Find(profile *Profile) ([]Email, error) {
	var emails []Email
	err := Mdb.db.Model(profile).Related(&emails)
	if err.Error != nil {
		Helpers.BaseLog("DATABASE", "ERROR", "", "NO ENTRY", 1062, 3, err.Error)
		return emails, err.Error
	}
	return emails, nil
}

//BeforeCreate Email
func (email *Email) BeforeCreate() {
	email.Crat = time.Now().Unix()
}

//Migrate Email
func (email *Email) Migrate() {
	Mdb.db.DropTable(Email{})
	Mdb.db.CreateTable(Email{})
}
