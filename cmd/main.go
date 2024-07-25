package main

import (
	"flag"
	"fmt"
	"pet-dex-backend/v2/pkg/migration"
)

func main() {

	upFlag := flag.Bool("up", false, "Run migrations UP")
	flag.Parse()

	if *upFlag {
		fmt.Println("Running Migrations UP...")
		migration.Up()
		fmt.Println("Migrations executed!")
		return
	}


	var number string
	fmt.Println("Migrations CLI")
	fmt.Println("Type the number of the command desired:\n1-Migrations UP\n2-Migrations DOWN\n3-Create a new migration")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Error while reading the values", err)
	}


	if number == "1" {
		fmt.Println("Running Migrations UP...")
		migration.Up()
		fmt.Println("Migrations executed!")
		return
	}

	if number == "2" {
		fmt.Println("Running Migrations DOWN...")
		migration.Down()
		fmt.Println("Migrations executed!")
		return
	}

	if number == "3" {
		fmt.Println("Type the name of the migration desired:")
		var name string
		_, err := fmt.Scan(&name)
		if err != nil {
			fmt.Println("Error while reading the values", err)
		}
		fmt.Println("Creating a new migration...")
		migration.Create(name)
		fmt.Println("Migration created!")
	}

}
