package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgresql://postgres.fdahxiqcopobqkoltzvl:ORKxf6aBAd4CeB6S@aws-0-ap-southeast-1.pooler.supabase.com:6543/postgres"

	fmt.Println("Trying to connect to database...")
	fmt.Println("Connection string:", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	fmt.Println("Database opened, testing ping...")

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}

	fmt.Println("✅ Database connection successful!")

	var version string
	err = db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		log.Fatal("Error querying database:", err)
	}

	fmt.Println("PostgreSQL version:", version)
}
