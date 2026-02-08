package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	// Load config
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	connStr := viper.GetString("DB_CONN")

	fmt.Println("Connecting to database...")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping:", err)
	}

	fmt.Println("✅ Database connected successfully!")

	// Run setup_complete.sql
	fmt.Println("\nRunning setup_complete.sql...")
	content, err := ioutil.ReadFile("database/setup_complete.sql")
	if err != nil {
		log.Fatal("Failed to read setup_complete.sql:", err)
	}

	_, err = db.Exec(string(content))
	if err != nil {
		log.Printf("Warning running setup_complete.sql: %v\n", err)
	} else {
		fmt.Println("✅ setup_complete.sql executed successfully!")
	}

	// Run migration_session3.sql
	fmt.Println("\nRunning migration_session3.sql...")
	content, err = ioutil.ReadFile("database/migration_session3.sql")
	if err != nil {
		log.Fatal("Failed to read migration_session3.sql:", err)
	}

	_, err = db.Exec(string(content))
	if err != nil {
		log.Printf("Warning running migration_session3.sql: %v\n", err)
	} else {
		fmt.Println("✅ migration_session3.sql executed successfully!")
	}

	fmt.Println("\n✨ All migrations completed!")
}
