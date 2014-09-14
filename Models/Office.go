package Models

import (
	"time"
	"github.com/devinceble/BaseServer/Helpers"
)

//TODO: Link to Constant Office From Organization Table
//TODO: Create Organization Table

//Office Model
type Office struct {
	Id        int64
	ProfileId int64
	Shortcode string `sql:"type:varchar(20);not null"`
	Office    string

	Crat int64
	Upat int64
	Dlat int64
}

//Find Office
func (office *Office) Find(profile *Profile) ([]Office, error){
	var offices []Office
	err := Mdb.db.Model(profile).Related(&offices)
	if err.Error != nil {
		Helpers.BaseLog("DATABASE", "ERROR", "", "NO ENTRY", 1062, 3, err.Error)
		return offices, err.Error
	}
	return offices, nil
}

//BeforeCreate Office
func (office *Office) BeforeCreate() {
	office.Crat = time.Now().Unix()
}

//Migrate Office
func (office *Office) Migrate() {
	Mdb.db.DropTable(Office{})
	Mdb.db.CreateTable(Office{})
}
