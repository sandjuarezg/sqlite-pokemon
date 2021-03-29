package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sandjuarezg/sqlite-pokemon/function"
)

type User struct {
	Id    int
	Name  string
	Pass  string
	Ocupa string
}

func AddUser(db *sql.DB) {
	var statement, err = db.Prepare("INSERT INTO users (name, password, ocupation) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	var user = User{}
	fmt.Println("Enter a name")
	fmt.Scan(&user.Name)
	fmt.Println("Enter a password")
	fmt.Scan(&user.Pass)

	var exit bool
	var opc int

	for !exit {
		fmt.Println("Select a ocupation")
		fmt.Println("1. Pokemon master")
		fmt.Println("2. Trainer")
		fmt.Println("3. Grabber")
		fmt.Println("4. Caretaker")
		fmt.Println("5. Traveler")
		fmt.Scan(&opc)

		switch opc {
		case 1:
			user.Ocupa = "Pokemon master"
			exit = true
		case 2:
			user.Ocupa = "Trainer"
			exit = true
		case 3:
			user.Ocupa = "Grabber"
			exit = true
		case 4:
			user.Ocupa = "Caretaker"
			exit = true
		case 5:
			user.Ocupa = "Traveler"
			exit = true
		default:
			fmt.Println("Option not valid")
			function.CleanConsole(2)
		}
	}
	statement.Exec(user.Name, user.Pass, user.Ocupa)
}

func ShowUser(db *sql.DB) {
	var rows, err = db.Query("SELECT id, name, password, ocupation FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var user = User{}
	fmt.Printf("|%-7s|%-15s|%-15s|%-15s|\n", "id", "Name", "Password", "Ocupation")
	fmt.Println("________________________________________________________")
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Pass, &user.Ocupa)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("|%-7d|%-15s|%-15s|%-15s|\n", user.Id, user.Name, user.Pass, user.Ocupa)
	}
}

func UpdateUser(db *sql.DB) (n int64) {
	var statement, err = db.Prepare("UPDATE users SET password = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	var user = User{}
	fmt.Println("Enter id")
	fmt.Scan(&user.Id)
	fmt.Println("Enter password to update")
	fmt.Scan(&user.Pass)

	var res, _ = statement.Exec(user.Pass, user.Id)
	n, _ = res.RowsAffected()

	return
}

func DeleteUser(db *sql.DB) (n int64) {
	var statement, err = db.Prepare("DELETE from users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	var user = User{}
	fmt.Println("Enter id")
	fmt.Scan(&user.Id)
	var res, _ = statement.Exec(user.Id)
	n, _ = res.RowsAffected()

	return
}

func SearchUser(db *sql.DB, id int) (user *User, err error) {
	var aux User
	var row = db.QueryRow("SELECT id, name, password, ocupation FROM users WHERE id = ?", id)
	err = row.Scan(&aux.Id, &aux.Name, &aux.Pass, &aux.Ocupa)
	if err != nil {
		return
	}
	user = &aux
	return
}
