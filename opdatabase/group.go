package opdatabase

import (
	"log"
)

//Groups for Different Products
//Can be creatd by user

//Group Struct
type Group struct {
	Id   int
	Name string
}

//CreateGroupTable creates a new Group table
func CreateGroupTable() {
	query := `CREATE TABLE gtable(
		id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(50) NOT NULL);`
	_, err := db.Exec(query)
	if err != nil {
		log.Println("Error in creating Group Table")
		log.Println(err)
	}
}

//CreateGroup create a new Group
func CreateGroup(g Group) {
	query := `INSERT INTO gtable (name) VALUES( ? );`
	_, err := db.Exec(query, g.Name)
	if err != nil {
		log.Println("error in creating an entry in Group table")
		log.Println(err)
	}
}

//SelectGroup selects all group
func SelectGroup() ([]Group, bool) {
	var g []Group
	query := `SELECT * FROM gtable;`
	r, err := db.Query(query)
	if err != nil {
		log.Print(err)
		return g, false
	}
	for r.Next() {
		var group Group
		err = r.Scan(&(group.Id), &(group.Name))
		if err != nil {
			log.Print(err)
			return g, false
		}
		g = append(g, group)
	}
	r.Close()
	return g, true
}

//EditGroup Edits a Group
func EditGroup(g Group) {
	query := `UPDATE gtable SET name= ? WHERE  id= ?`
	_, err := db.Exec(query, g.Name, g.Id)
	if err != nil {
		log.Println("Error Updating  an Entry in Group Table")
		log.Println(err)
	}
}

//DeleteGroup Deletes a Group
func DeleteGroup(g Group) {
	query := `DELETE FROM gtable WHERE id= ?; `
	_, err := db.Exec(query, g.Id)
	if err != nil {
		log.Println("Error iin deleting an entry in Group Table")
		log.Println(err)
	}
}
