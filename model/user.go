package model

import (
	"github.com/achhapolia10/inventory-manager/opdatabase"
)

//CreateUserTable wrapper for  opdatabase
func CreateUserTable() {
	opdatabase.CreateUserTable()
}

//CreateUser creates a user
func CreateUser(uname, p string, admin int) {
	opdatabase.CreateUser(uname, p, admin)
}

//GetUsers Return all users
func GetUsers() []opdatabase.User {
	u := opdatabase.SelectUsers()
	return u
}

//CheckUser checks a user
func CheckUser(uname, p string) (int, bool) {
	if uname != "" && p != "" {
		pass, admin, err := opdatabase.GetUser(uname)
		if err != nil {
			return 0, false
		}
		if p != pass {
			return 0, false
		}
		return admin, true

	}
	return 0, false
}

//UpdatePassword updates the password
func UpdatePassword(u, p string) {
	opdatabase.UpdateUser(u, p)
}

//DeleteUser deletes the user
func DeleteUser(u string) {
	opdatabase.DeleteUser(u)
}
