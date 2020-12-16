package models

import (
	"fmt"
	"time"

	"database/sql"

	// Postgres database driver
	_ "github.com/lib/pq"
)

// Database Represents the exported database connection obect
var Database *sql.DB

// Connect -> connects the application to the database
func Connect(host string, port string, dbName string, user string, pass string, maxTries int) {
	// Generate connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbName)

	// Try to open a connection
	dbCandidate, err := sql.Open("postgres", psqlInfo)

	// Send panic error if connection failed
	if err != nil {
		fmt.Println("Database connection params error")
		panic(err)
	}

	// Ping database
	err = dbCandidate.Ping()
	tests := 0

	// Retry if failed until max tries reached
	for err != nil && tests < maxTries {
		fmt.Println("Connection to database did not succeed, new try")

		// Sleep 10 seconds before new retry
		time.Sleep(10 * time.Second)
		dbCandidate, err = sql.Open("postgres", psqlInfo)
		err = dbCandidate.Ping()

		tests++
	}

	if err != nil {
		// Max tries reached and connection still failing
		fmt.Println("Database initialisation error")
		panic(err)
	}

	// Once connection succedded, set Database global variable with candidated one
	Database = dbCandidate

	// Inform successfull connection
	fmt.Println("Database connected successfully!")
}
