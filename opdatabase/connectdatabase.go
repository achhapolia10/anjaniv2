package opdatabase

import (
	"database/sql"
	"fmt"
)

//ConnectDatabase connects to database Server at the start of the server
func ConnectDatabase() {
	fmt.Println("Connecting to the SQL server")
	var err error
	db, err = sql.Open("mysql", "root:ilijksms1999@/anjani_test")
	err1 := db.Ping()
	if err1 != nil {
		fmt.Println(err)
	}
	fmt.Println("Database Server connected")

}
