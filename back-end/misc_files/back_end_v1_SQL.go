package misc_files

//import (
//	"database/sql"
//	"fmt"
//
//	_ "github.com/mattn/go-sqlite3"
//)
//
//func main() {
//	// Open a connection to the SQLite database
//	db, err := sql.Open("sqlite3", "./test.db")
//	if err != nil {
//		fmt.Println("Error opening database:", err)
//		return
//	}
//	defer db.Close()
//
//	// Create a table
//	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)`)
//	if err != nil {
//		fmt.Println("Error creating table:", err)
//		return
//	}
//
//	// Insert a row of data
//	_, err = db.Exec(`INSERT INTO users (name) VALUES (?)`, "John Doe")
//	if err != nil {
//		fmt.Println("Error inserting data:", err)
//		return
//	}
//
//	// Query the data
//	rows, err := db.Query(`SELECT id, name FROM users`)
//	if err != nil {
//		fmt.Println("Error querying data:", err)
//		return
//	}
//	defer rows.Close()
//
//	// Print the data
//	for rows.Next() {
//		var id int
//		var name string
//		err := rows.Scan(&id, &name)
//		if err != nil {
//			fmt.Println("Error scanning rows:", err)
//			return
//		}
//		fmt.Printf("%d: %s\n", id, name)
//	}
//}
