package Models

import "time"

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

//BeforeCreate Profile
func (address *Address) BeforeCreate() {
	address.Crat = time.Now().Unix()
}

//Migrate User
func (address *Address) Migrate() {
	Mdb.db.DropTable(Address{})
	Mdb.db.CreateTable(Address{})
}
