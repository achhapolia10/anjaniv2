package opdatabase

import (
	"database/sql"
	"fmt"
	"log"
)

//ConnectDatabase connects to database Server at the start of the server
func ConnectDatabase(name string) {
	fmt.Println("Connecting to the SQL server")
	var err error
	db, err = sql.Open("sqlite3", name)
	err1 := db.Ping()
	if err1 != nil {
		log.Println(err)
	}
	fmt.Println("Database Server connected")

}
