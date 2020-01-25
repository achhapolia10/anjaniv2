package opdatabase

import "log"

//User Defines the User Type
type User struct {
	ID         int
	Username   string
	Admin      int
	IsLoggedIn bool
	Password   string
}

//CreateUserTable creates the user table
func CreateUserTable() {
	query := `CREATE TABLE user(
		id 	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(50),
		password VARCHAR(50),
		admin	INTEGER );`
	_, err := db.Exec(query)
	if err != nil {
		log.Printf("Error in Creating User Table: %v", err)
		return
	}
	CreateUser("admin", "ilijksms1999", 1)
}

//CreateUser Creates an User
func CreateUser(uname, p string, admin int) {
	query := `INSERT INTO user (username,password,admin) VALUES(?,?,?);`
	_, err := db.Exec(query, uname, p, admin)
	if err != nil {
		log.Println("error in creating an entry in Group table")
		log.Println(err)
	}

}

//GetUser gets User
func GetUser(uname string) (string, int, error) {
	var id, admin int
	var u, p string
	query := `SELECT * FROM user WHERE username=?`
	r, err := db.Query(query, uname)
	if err != nil {
		log.Printf("Error in reading user table: %v", err)
		return "", 3, err
	}
	if r.Next() {
		err := r.Scan(&id, &u, &p, &admin)
		if err != nil {
			log.Printf("Error in reading user table: %v", err)
		}
	}
	r.Close()
	return p, admin, nil
}

//SelectUsers gets the Users
func SelectUsers() []User {
	query := `SELECT * FROM user;`
	var users []User
	r, err := db.Query(query)
	if err != nil {
		log.Printf("Error in reading user table: %v", err)
	}
	for r.Next() {
		var user User
		err := r.Scan(&(user.ID), &(user.Username), &(user.Password), &(user.Admin))
		if err != nil {
			log.Printf("Error in reading user table: %v", err)
			break
		}
		users = append(users, user)
	}
	return users
}

//UpdateUser Updates the password
func UpdateUser(uname, p string) {
	query := `UPDATE user SET password = ? WHERE username=?`
	_, err := db.Exec(query, p, uname)
	if err != nil {
		log.Printf("Error in Creating User Table: %v", err)
	}
}

//DeleteUser Deletes the Username
func DeleteUser(uname string) {
	query := `DELETE FROM user WHERE username=?`
	_, err := db.Exec(query, uname)
	if err != nil {
		log.Printf("Error in Creating User Table: %v", err)
	}

}
