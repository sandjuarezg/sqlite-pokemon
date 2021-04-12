package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sandjuarezg/sqlite-pokemon/function"
)

type User struct {
	Id    int
	Name  string
	Pass  string
	Ocupa string
}

func AddUser(db *sql.DB) (err error) {
	smt, err := db.Prepare("INSERT INTO users (name, password, ocupation) VALUES (?, ?, ?)")
	if err != nil {
		err = function.ErrInsert
		return
	}
	defer smt.Close()

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
	_, err = smt.Exec(user.Name, user.Pass, user.Ocupa)
	if err != nil {
		err = function.ErrInsert
		return
	}

	return
}

func ShowUser(db *sql.DB) (err error) {
	rows, err := db.Query("SELECT id, name, password, ocupation FROM users")
	if err != nil {
		err = function.ErrShowData
		return
	}
	defer rows.Close()

	var user = User{}
	fmt.Printf("|%-7s|%-15s|%-15s|%-15s|\n", "id", "Name", "Password", "Ocupation")
	fmt.Println("________________________________________________________")
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Pass, &user.Ocupa)
		if err != nil {
			err = function.ErrScan
			return
		}
		fmt.Printf("|%-7d|%-15s|%-15s|%-15s|\n", user.Id, user.Name, user.Pass, user.Ocupa)
	}

	return
}

func UpdateUser(db *sql.DB) (n int64, err error) {
	var user = User{}
	fmt.Println("Enter id")
	fmt.Scan(&user.Id)
	fmt.Println("Enter password to update")
	fmt.Scan(&user.Pass)

	row, err := db.Exec("UPDATE users SET password = ? WHERE id = ?", user.Pass, user.Id)
	if err != nil {
		err = function.ErrUpdate
		return
	}
	n, err = row.RowsAffected()
	if err != nil {
		err = function.ErrUpdate
		return
	}

	return
}

func DeleteUser(db *sql.DB) (n int64, err error) {
	var user = User{}
	fmt.Println("Enter id")
	fmt.Scan(&user.Id)

	row, err := db.Exec("DELETE from users WHERE id = ?", user.Id)
	if err != nil {
		err = function.ErrDelete
		return
	}
	n, err = row.RowsAffected()
	if err != nil {
		err = function.ErrDelete
		return
	}

	return
}

func SearchUser(db *sql.DB, id int) (user *User, err error) {
	user = new(User)
	var row = db.QueryRow("SELECT id FROM users WHERE id = ?", id)
	err = row.Scan(&user.Id)
	if err != nil {
		err = function.ErrScan
		return
	}
	return
}
