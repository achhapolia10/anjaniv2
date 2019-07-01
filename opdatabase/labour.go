package opdatabase

import "log"

//CreateLabourTable creates the table
func CreateLabourTable() {
	query := `CREATE TABLE labour(
		name VARCHAR(50) NOT NULL,
		date VARCHAR(50) NOT NULL);`
	_, err := db.Exec(query)
	if err != nil {
		log.Println("Error in Creating Labour Table:", err)
	}
}

//AddLabour adds an Entry to labour table
func AddLabour(n, d string) {
	query := `INSERT INTO labour (name,date) VALUES(?,?)`

	if _, err := db.Exec(query, n, d); err != nil {
		log.Println("Error in instert the labour names : ", err)
	}
}

//SelectLabours gives all the labour names
func SelectLabours() map[string]bool {

	names := make(map[string]bool)
	query := `SELECT * FROM labour;`
	r, err := db.Query(query)

	if err != nil {
		log.Println("Error in Reading labour Names: ", err)
	}

	for r.Next() {
		var name, d string
		err = r.Scan(&name, &d)
		if err != nil {
			log.Println("Error in Creating Labour Table:", err)
			break
		}
		names[name] = true
	}
	return names
}

//UpdateLabour updates the Date for a labour
func UpdateLabour(n, d string) bool {

	query := `UPDATE labour SETE date=? WHERE name=?`

	if _, err := db.Exec(query, d); err != nil {
		log.Println("Error in updating the labour names : ", err)
	}
	return false
}

//DeleteLabours delete all the labour for a date
func DeleteLabours(d string) {
	query := `DELETE FROM labour WHERE date=?`

	if _, err := db.Exec(query, d); err != nil {
		log.Println("Error in deleting the labour names : ", err)
	}
}
