package opdatabase

import "log"

//CreateUserTable creates the user table
func CreateUserTable() {
	query := `CREATE TABLE user(
		id 	INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
		username VARCHAR(50),
		password VARCHAR(50),
		admin	INT);`
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
	return p, admin, nil
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
