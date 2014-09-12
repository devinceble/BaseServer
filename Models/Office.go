package Models

import "time"

//TODO: Link to COnstant Office From Organization Table
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

func init() {
	Mdb.db, _ = Connect()
}

//BeforeCreate Profile
func (office *Office) BeforeCreate() {
	office.Crat = time.Now().Unix()
}

//Migrate User
func (office *Office) Migrate() {
	Mdb.db.DropTable(Office{})
	Mdb.db.CreateTable(Office{})
}
