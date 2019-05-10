package opdatabase

import "log"

//Groups for Different Products
//Can be creatd by user

//Group Struct
type Group struct {
	id   int
	name string
}

//CreateGroupTable creates a new Group table
func CreateGroupTable() {
	query := `CREATE TABLE group(
		id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(50) NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Println("Error in creating Group Table")
		log.Println(err)
	}
}

//CreateGroup create a new Group
func CreateGroup(g Group) {
	query := `INSERT INTO group (name) VALUES( ? );`
	_, err := db.Exec(query, g.name)
	if err != nil {
		log.Println("error in creating an entry in Group table")
		log.Println(err)
	}
}

//EditGroup Edits a Group
func EditGroup(g Group) {
	query := `UPDATE group SET name= ? WHERE  id= ?`
	_, err := db.Exec(query, g.name, g.id)
	if err != nil {
		log.Println("Error Updating  an Entry in Group Table")
		log.Println(err)
	}
}

//DeleteGroup Deletes a Group
func DeleteGroup(g Group) {
	query := `DELETE FROM group WHERE id= ?; `
	_, err := db.Exec(query, g.id)
	if err != nil {
		log.Println("Error iin deleting an entry in Group Table")
		log.Println(err)
	}
}
