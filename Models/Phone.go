package Models

import "time"

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

func init() {
	Mdb.db, _ = Connect()
}

//BeforeCreate Profile
func (phone *Phone) BeforeCreate() {
	phone.Crat = time.Now().Unix()
}

//Migrate User
func (phone *Phone) Migrate() {
	Mdb.db.DropTable(Phone{})
	Mdb.db.CreateTable(Phone{})
}
