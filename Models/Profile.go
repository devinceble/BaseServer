package Models

import "time"

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

func init() {
	Mdb.db, _ = Connect()
}

//BeforeCreate Profile
func (profile *Profile) BeforeCreate() {
	profile.Crat = time.Now().Unix()
}

//Migrate User
func (profile *Profile) Migrate() {
	Mdb.db.DropTable(Profile{})
	Mdb.db.CreateTable(Profile{})
}
