package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sandjuarezg/sqlite-pokemon/function"
)

type Attack struct {
	Id      int
	Name    string
	Power   int
	Defense int
	Speed   int
}

func AddAttack(db *sql.DB) (err error) {
	statement, err := db.Prepare("INSERT INTO attacks (name, power, defense, speed) VALUES (?, ?, ?, ?)")
	if err != nil {
		err = function.ErrInsert
		return
	}
	defer statement.Close()

	var attacks = Attack{}
	fmt.Println("Enter a name")
	fmt.Scan(&attacks.Name)
	fmt.Println("Enter power")
	fmt.Scan(&attacks.Power)
	fmt.Println("Enter defense")
	fmt.Scan(&attacks.Defense)
	fmt.Println("Enter speed")
	fmt.Scan(&attacks.Speed)

	statement.Exec(attacks.Name, attacks.Power, attacks.Defense, attacks.Speed)

	return
}

func ShowAttacks(db *sql.DB) (err error) {
	rows, err := db.Query("SELECT id, name, power, defense, speed FROM attacks")
	if err != nil {
		err = function.ErrShowData
		return
	}
	defer rows.Close()

	var attacks = Attack{}
	fmt.Printf("|%-9s|%-15s|%-15s|%-15s|%-15s|\n", "id", "Name", "Power", "defense", "Speed")
	fmt.Println("___________________________________________________________________________")
	for rows.Next() {
		err = rows.Scan(&attacks.Id, &attacks.Name, &attacks.Power, &attacks.Defense, &attacks.Speed)
		if err != nil {
			err = function.ErrScan
			return
		}
		fmt.Printf("|%-9d|%-15s|%-15d|%-15d|%-15d|\n", attacks.Id, attacks.Name, attacks.Power, attacks.Defense, attacks.Speed)
	}

	return
}

func UpdateAttacks(db *sql.DB) (n int64, err error) {
	statement, err := db.Prepare("UPDATE attacks SET name = ? WHERE id = ?")
	if err != nil {
		err = function.ErrUpdate
		return
	}
	defer statement.Close()

	var attacks = Attack{}
	fmt.Println("Enter id")
	fmt.Scan(&attacks.Id)
	fmt.Println("Enter name to update")
	fmt.Scan(&attacks.Name)

	var res, _ = statement.Exec(attacks.Name, attacks.Id)
	n, _ = res.RowsAffected()

	return
}

func DeleteAttacks(db *sql.DB) (n int64, err error) {
	statement, err := db.Prepare("DELETE from attacks WHERE id = ?")
	if err != nil {
		err = function.ErrDelete
		return
	}
	defer statement.Close()

	var attacks = Attack{}
	fmt.Println("Enter id")
	fmt.Scan(&attacks.Id)
	var res, _ = statement.Exec(attacks.Id)
	n, _ = res.RowsAffected()

	return
}

func SearchAttacks(db *sql.DB, id int) (attacks *Attack, err error) {
	var aux Attack
	var row = db.QueryRow("SELECT id, name, power, defense, speed FROM attacks WHERE id = ?", id)
	err = row.Scan(&aux.Id, &aux.Name, &aux.Power, &aux.Defense, &aux.Speed)
	if err != nil {
		err = function.ErrScan
		return
	}
	attacks = &aux
	return
}
