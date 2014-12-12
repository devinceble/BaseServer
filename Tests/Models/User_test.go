package TestsModels

import (
	. "github.com/devinceble/BaseServer/Models"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
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

		user.Migrate()
		profile.Migrate()
		address.Migrate()
		office.Migrate()
		email.Migrate()
		phone.Migrate()

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
		Convey("Finding User", func() {
			var UserId int64 = 1
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

		Convey("Find All User", func() {
			users, _ := user.FindAll()
			So(len(users), ShouldBeGreaterThan, 0)
		})

		Convey("Updating User", func() {

		})

		Convey("Deleting User", func() {

		})

		Convey("Restoring User", func() {

		})
	})

}
