package migration

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"pet-dex-backend/v2/infra/config"
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
