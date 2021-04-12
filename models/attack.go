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
	smt, err := db.Prepare("INSERT INTO attacks (name, power, defense, speed) VALUES (?, ?, ?, ?)")
	if err != nil {
		err = function.ErrInsert
		return
	}
	defer smt.Close()

	var attacks = Attack{}
	fmt.Println("Enter a name")
	fmt.Scan(&attacks.Name)
	fmt.Println("Enter power")
	fmt.Scan(&attacks.Power)
	fmt.Println("Enter defense")
	fmt.Scan(&attacks.Defense)
	fmt.Println("Enter speed")
	fmt.Scan(&attacks.Speed)

	_, err = smt.Exec(attacks.Name, attacks.Power, attacks.Defense, attacks.Speed)
	if err != nil {
		err = function.ErrInsert
		return
	}

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
	var attacks = Attack{}
	fmt.Println("Enter id")
	fmt.Scan(&attacks.Id)
	fmt.Println("Enter name to update")
	fmt.Scan(&attacks.Name)

	row, err := db.Exec("UPDATE attacks SET name = ? WHERE id = ?", attacks.Name, attacks.Id)
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

func DeleteAttacks(db *sql.DB) (n int64, err error) {
	var attacks = Attack{}
	fmt.Println("Enter id")
	fmt.Scan(&attacks.Id)

	row, err := db.Exec("DELETE from attacks WHERE id = ?", attacks.Id)
	if err != nil {
		err = function.ErrShowData
		return
	}
	n, err = row.RowsAffected()
	if err != nil {
		err = function.ErrUpdate
		return
	}

	return
}

func SearchAttacks(db *sql.DB, id int) (attacks *Attack, err error) {
	attacks = new(Attack)
	var row = db.QueryRow("SELECT id FROM attacks WHERE id = ?", id)
	err = row.Scan(&attacks.Id)
	if err != nil {
		return
	}
	return
}
