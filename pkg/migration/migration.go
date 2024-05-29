package migration

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"pet-dex-backend/v2/infra/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Up() {
	env, err := config.LoadEnv("../")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v\n", err)
	}
	fmt.Println(env.DBUrl_Migration)
	db, err := sql.Open("mysql", env.DBUrl_Migration)
	if err != nil {
		log.Fatalf("Failed connecting to the database: %v\n", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	migration, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}
	err = migration.Up()
	if err != nil {
		log.Fatalf("Failed on running migrations up: %v\n", err)
		return
	}
}

func Down() {
	env, err := config.LoadEnv("../")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v\n", err)
	}
	fmt.Println(env.DBUrl_Migration)
	db, err := sql.Open("mysql", env.DBUrl_Migration)
	if err != nil {
		log.Fatalf("Failed connecting to the database: %v\n", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	migration, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}
	err = migration.Down()
	if err != nil {
		log.Fatalf("Failed on running migrations down: %v\n", err)
		return
	}

}

func Create(name string) {
	path, err := config.LoadEnv("../../")
	if err != nil {
		fmt.Println("Error loading the .env file:", err)
	}
	data := time.Now()
	timestamp := data.Format("20060102150405")
	fmt.Println("Current date and time: ", timestamp)
	fileNameDown := fmt.Sprintf("%s/%s_%s.down.sql", path.MIGRATIONS_PATH, timestamp, name)
	fileNameUp := fmt.Sprintf("%s/%s_%s.up.sql", path.MIGRATIONS_PATH, timestamp, name)
	// Create the file
	fileDown, err := os.Create(fileNameDown)
	if err != nil {
		fmt.Println("Error creating down file:", err)
		return
	}
	defer fileDown.Close()

	fileUp, err := os.Create(fileNameUp)
	if err != nil {
		fmt.Println("Error creating up file:", err)
		return
	}
	defer fileUp.Close()
}
