package Models

import (
	"time"
	"github.com/devinceble/BaseServer/Helpers"
)

//Address Model
type Address struct {
	Id        int64
	ProfileId int64
	Type      string `sql:"type:varchar(20);not null"`
	House     string
	Street    string
	Barangay  string
	Zone      string
	City      string
	Province  string
	Zipcode   int64

	Crat int64
	Upat int64
	Dlat int64
}

//Find Address
func (address *Address) Find(profile *Profile) ([]Address, error){
	var addresses []Address
	err := Mdb.db.Model(profile).Related(&addresses)
	if err.Error != nil {
		Helpers.BaseLog("DATABASE", "ERROR", "", "NO ENTRY", 1062, 3, err.Error)
		return addresses, err.Error
	}
	return addresses, nil
}

//BeforeCreate Address
func (address *Address) BeforeCreate() {
	address.Crat = time.Now().Unix()
}

//Migrate Address
func (address *Address) Migrate() {
	Mdb.db.DropTable(Address{})
	Mdb.db.CreateTable(Address{})
}
