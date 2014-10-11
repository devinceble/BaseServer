package TestsModels

import (
	"testing"
	. "github.com/devinceble/BaseServer/Models"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_UserModel(t *testing.T) {
	Convey("Creating User", t, func() {
		var user User
		var profile Profile
		var address Address
		var oaddress Address
		var office Office
		var email Email
		var phone Phone

		var user2 User
		var profile2 Profile
		var address2 Address
		var office2 Office
		var email2 Email
		var phone2 Phone

		user.Migrate()
		profile.Migrate()
		address.Migrate()
		office.Migrate()
		email.Migrate()
		phone.Migrate()

		profile2.FirstName = "Maria Del Rie"
		profile2.LastName = "Marte"
		profile2.MiddleName = "Sabana"

		address2.Type = "Home"
		address2.House = "01-78"
		address2.Street = "Magsaysay"
		address2.Barangay = "2"
		address2.Zone = "1"
		address2.City = "Malaybalay"
		address2.Province = "Bukidnon"
		address2.Zipcode = 8700

		office2.Shortcode = "PG"
		office2.Office = "Pryce Garden"

		email2.Type = "Personal"
		email2.Email = "mariadelrie@gmail.com"

		phone2.Type = "Personal"
		phone2.Phone = "+639180000001"

		user2.Username = "mariadelrie"
		user2.Password = "qwerty"
		user2.Profile = profile2
		user2.Profile.Address = []Address{address2}
		user2.Profile.Office = []Office{office2}
		user2.Profile.Email = []Email{email2}
		user2.Profile.Phone = []Phone{phone2}

		err1 := user2.Create()
		So(err1, ShouldBeNil)

		profile.FirstName = "Leivince John"
		profile.LastName = "Marte"
		profile.MiddleName = "Diez"

		address.Type = "Home"
		address.House = "01-78"
		address.Street = "Magsaysay"
		address.Barangay = "2"
		address.Zone = "1"
		address.City = "Malaybalay"
		address.Province = "Bukidnon"
		address.Zipcode = 8700

		oaddress.Type = "Office"
		oaddress.House = "00-00"
		oaddress.Street = "Capitol Compound"
		oaddress.Barangay = "1"
		oaddress.Zone = "7"
		oaddress.City = "Malaybalay"
		oaddress.Province = "Bukidnon"
		oaddress.Zipcode = 8700

		office.Shortcode = "PAO"
		office.Office = "Provincial Accountants Office"

		email.Type = "Personal"
		email.Email = "devinceble@gmail.com"

		phone.Type = "Personal"
		phone.Phone = "+639180000000"

		user.Username = "devinceble"
		user.Password = "qwerty"
		user.Profile = profile
		user.Profile.Address = []Address{address, oaddress}
		user.Profile.Office = []Office{office}
		user.Profile.Email = []Email{email}
		user.Profile.Phone = []Phone{phone}

		err := user.Create()
		So(err, ShouldBeNil)
		So(user.Upat, ShouldBeZeroValue)
		So(user.Dlat, ShouldBeZeroValue)

		Convey("User Profile", func() {
			So(user.Id, ShouldEqual, user.Profile.UserId)
			So(user.Profile.Crat, ShouldBeGreaterThan, 0)
			So(user.Profile.Upat, ShouldBeZeroValue)
			So(user.Profile.Dlat, ShouldBeZeroValue)

			Convey("User Profile Address", func() {
				So(user.Profile.Id, ShouldEqual, user.Profile.Address[0].ProfileId)
				So(user.Profile.Address[0].Crat, ShouldBeGreaterThan, 0)
				So(user.Profile.Address[0].Upat, ShouldBeZeroValue)
				So(user.Profile.Address[0].Dlat, ShouldBeZeroValue)
			})

			Convey("User Profile Office", func() {
				So(user.Profile.Id, ShouldEqual, user.Profile.Office[0].ProfileId)
				So(user.Profile.Office[0].Crat, ShouldBeGreaterThan, 0)
				So(user.Profile.Office[0].Upat, ShouldBeZeroValue)
				So(user.Profile.Office[0].Dlat, ShouldBeZeroValue)
			})

			Convey("User Profile Email", func() {
				So(user.Profile.Id, ShouldEqual, user.Profile.Email[0].ProfileId)
				So(user.Profile.Email[0].Crat, ShouldBeGreaterThan, 0)
				So(user.Profile.Email[0].Upat, ShouldBeZeroValue)
				So(user.Profile.Email[0].Dlat, ShouldBeZeroValue)
			})

			Convey("User Profile Phone", func() {
				So(user.Profile.Id, ShouldEqual, user.Profile.Phone[0].ProfileId)
				So(user.Profile.Phone[0].Crat, ShouldBeGreaterThan, 0)
				So(user.Profile.Phone[0].Upat, ShouldBeZeroValue)
				So(user.Profile.Phone[0].Dlat, ShouldBeZeroValue)
			})
		})
	})

	Convey("Finding User", t, func() {
		var user User
		var UserId int64 = 2
		user.Find(UserId)
		So(user.Id, ShouldEqual, UserId)
		Convey("Get User Profile", func() {
			So(user.Profile.UserId, ShouldEqual, UserId)

			Convey("Get User Profile Address", func() {
				So(user.Profile.Address[0].ProfileId, ShouldEqual, user.Profile.Id)
			})

			Convey("Get User Profile Office", func() {
				So(user.Profile.Office[0].ProfileId, ShouldEqual, user.Profile.Id)
			})

			Convey("Get User Profile Email", func() {
				So(user.Profile.Email[0].ProfileId, ShouldEqual, user.Profile.Id)
			})

			Convey("Get User Profile Phone", func() {
				So(user.Profile.Phone[0].ProfileId, ShouldEqual, user.Profile.Id)
			})

		})
	})

	Convey("Find All User", t, func() {
		var user User
		users, _ := user.FindAll()
		So(len(users), ShouldBeGreaterThan, 0)
	})

	Convey("Updating User", t, func() {

	})

	Convey("Deleting User", t, func() {

	})

	Convey("Restoring User", t, func() {

	})

}
