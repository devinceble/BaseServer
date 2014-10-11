package Models

import (
	"time"

	"github.com/devinceble/BaseServer/Helpers"
)

//Profile Model
type Profile struct {
	Id            int64
	UserId        int64
	FirstName     string `sql:"not null"`
	LastName      string `sql:"not null"`
	MiddleName    string `sql:"not null"`
	ExtensionName string

	Crat int64
	Upat int64
	Dlat int64

	Address []Address
	Office  []Office
	Email   []Email
	Phone   []Phone
}

//Find Profile
func (profile *Profile) Find(user *User) (*Profile, error) {
	err := Mdb.db.Model(user).Related(profile)
	if err.Error != nil {
		Helpers.BaseLog("DATABASE", "ERROR", "", "NO ENTRY", 1062, 3, err.Error)
		return profile, err.Error
	}
	return profile, nil
}

//BeforeCreate Profile
func (profile *Profile) BeforeCreate() {
	profile.Crat = time.Now().Unix()
}

//AfterFind Profile
func (profile *Profile) AfterFind() {
	var address Address
	addss, _ := address.Find(profile)
	profile.Address = addss
	var ofs Office
	ofss, _ := ofs.Find(profile)
	profile.Office = ofss
	var eml Email
	emls, _ := eml.Find(profile)
	profile.Email = emls
	var phn Phone
	phns, _ := phn.Find(profile)
	profile.Phone = phns
}

//Migrate Profile
func (profile *Profile) Migrate() {
	Mdb.db.DropTable(Profile{})
	Mdb.db.CreateTable(Profile{})
}
