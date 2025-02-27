package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "go_postgres_demo"
)

var db *sql.DB

type User struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database! %v", err))
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("Failed to ping database! %v", err))
	}
	fmt.Println("Successfully connected to database!")

	//addNewRecord()
	//updateRecord()
	//fetchRecord(3)
	//deleteRecord(6)
	fetchAllRecords()

}

func addNewRecord() {
	sqlStatement := `
	INSERT INTO users (age, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, 5, "hana@sol.com", "hana", "soltani").Scan(&id)
	if err != nil {
		panic(fmt.Sprintf("Failed to insert a row! %v", err))
	}
	fmt.Println("New record ID is:", id)
}

func updateRecord() {
	sqlStatement := `
	UPDATE users
	SET age = $2
	WHERE id = $1`
	result, err := db.Exec(sqlStatement, 4, 4)
	if err != nil {
		panic(fmt.Sprintf("Failed to update a row! %v", err))
	}
	count, err := result.RowsAffected()
	if err != nil {
		panic(fmt.Sprintf("Failed to get rows affected! %v", err))
	}

	fmt.Printf("%v rows updated. Successfully updated the record!\n", count)
}

func fetchRecord(id int) {
	sqlStatement := `SELECT id, first_name, last_name, email FROM users WHERE id=$1;`
	//var firstName, lastName, email string
	var user User
	err := db.QueryRow(sqlStatement, id).Scan(&id, &user.FirstName, &user.LastName, &user.Email)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id, user.FirstName, user.LastName, user.Email)
	default:
		panic(fmt.Sprintf("Failed to fetch a row! %v", err))
	}
}

func fetchAllRecords() {
	row, err := db.Query("SELECT id, first_name, last_name, email FROM users;")
	if err != nil {
		panic(fmt.Sprintf("Failed to fetch rows! %v", err))
	}
	defer row.Close()

	for row.Next() {
		var id int
		var firstName, lastName, email string
		err := row.Scan(&id, &firstName, &lastName, &email)
		if err != nil {
			panic(fmt.Sprintf("Failed to scan a row! %v", err))
		}
		fmt.Println(id, firstName, lastName, email)
	}

	err = row.Err()
	if err != nil {
		panic(fmt.Sprintf("Failed to fetch all rows! %v", err))
	}
}

func deleteRecord(id int) {
	sqlStatement := `DELETE FROM users WHERE id=$1;`
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		panic(fmt.Sprintf("Failed to delete a row! %v", err))
	}
	fmt.Println("Successfully deleted the record!")
}
