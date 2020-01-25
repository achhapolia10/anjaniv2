package opdatabase

import (
	"database/sql"
	"fmt"
	"log"
)

//ConnectDatabase connects to database Server at the start of the server
func ConnectDatabase() {
	fmt.Println("Connecting to the SQL server")
	var err error
	db, err = sql.Open("sqlite3", "./test.db")
	err1 := db.Ping()
	if err1 != nil {
		log.Println(err)
	}
	fmt.Println("Database Server connected")

}
