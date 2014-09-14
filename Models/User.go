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

//Find UserById
func (user *User) Find(Id int64) (*User, error){
	err := Mdb.db.Select("id, username").First(user, Id)
	if err.Error != nil {
		Helpers.BaseLog("DATABASE", "ERROR", "", "NO ENTRY", 1062, 3, err.Error)
		return user, err.Error
	}
	return user, nil
}

//Find UserAll
func (user *User) FindAll() ([]User, error){
	var users []User
	err := Mdb.db.Find(&users)
	if err.Error != nil {
		Helpers.BaseLog("DATABASE", "ERROR", "", "NO ENTRY", 1062, 3, err.Error)
		return users, err.Error
	}
	return users, nil
}


//BeforeCreate User
func (user *User) BeforeCreate() {
	user.Crat = time.Now().Unix()
}

//AfterFind User
func (user *User) AfterFind() {
	user.Profile.Find(user)
}

//Migrate User
func (user *User) Migrate() {
	Mdb.db.DropTable(User{})
	Mdb.db.CreateTable(User{})
}
