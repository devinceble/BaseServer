package Models

import (
	"time"
	"github.com/devinceble/BaseServer/Helpers"
)

//TODO: GET Standard Phone Structure

//Phone Model
type Phone struct {
	Id        int64
	ProfileId int64
	Type      string `sql:"type:varchar(20);not null"`
	Phone     string

	Crat int64
	Upat int64
	Dlat int64
}
//Find Phone
func (phone *Phone) Find(profile *Profile) ([]Phone, error){
	var phones []Phone
	err := Mdb.db.Model(profile).Related(&phones)
	if err.Error != nil {
		Helpers.BaseLog("DATABASE", "ERROR", "", "NO ENTRY", 1062, 3, err.Error)
		return phones, err.Error
	}
	return phones, nil
}
//BeforeCreate Phone
func (phone *Phone) BeforeCreate() {
	phone.Crat = time.Now().Unix()
}

//Migrate Phone
func (phone *Phone) Migrate() {
	Mdb.db.DropTable(Phone{})
	Mdb.db.CreateTable(Phone{})
}
