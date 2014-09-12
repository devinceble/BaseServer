package Models

import (
	"time"

	"github.com/devinceble/BaseServer/Helpers"
	"github.com/jameskeane/bcrypt"
)

//User Model
type User struct {
	Id       int64
	Username string `sql:"type:varchar(20);not null;unique"`
	Password string `sql:"-"`
	HashByte string `sql:"not null;unique"`

	Crat int64
	Upat int64
	Dlat int64

	Profile Profile
}

func init() {
	Mdb.db, _ = Connect()
}

//Create User
func (user *User) Create() error {
	salt, _ := bcrypt.Salt(10)
	user.HashByte, _ = bcrypt.Hash(user.Password, salt)
	err := Mdb.db.Create(user)
	if err.Error != nil {
		Helpers.BaseLog("DATABASE", "ERROR", "", "DUPLICATE", 1062, 3, err.Error)
		return err.Error
	}
	return nil
}

//BeforeCreate User
func (user *User) BeforeCreate() {
	user.Crat = time.Now().Unix()
}

//Migrate User
func (user *User) Migrate() {
	Mdb.db.DropTable(User{})
	Mdb.db.CreateTable(User{})
}
