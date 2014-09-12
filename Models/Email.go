package Models

import "time"

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

func init() {
	Mdb.db, _ = Connect()
}

//BeforeCreate Email
func (email *Email) BeforeCreate() {
	email.Crat = time.Now().Unix()
}

//Migrate User
func (email *Email) Migrate() {
	Mdb.db.DropTable(Email{})
	Mdb.db.CreateTable(Email{})
}
