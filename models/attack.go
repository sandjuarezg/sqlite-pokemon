package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Attack struct {
	Id      int
	Name    string
	Power   int
	Defense int
	Speed   int
}

func AddAttacks(db *sql.DB) {
	statement, err := db.Prepare("INSERT INTO attacks (name, power, defense, speed) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
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
}

func ShowAttacks(db *sql.DB) {
	var rows, err = db.Query("SELECT id, name, power, defense, speed FROM attacks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var attacks = Attack{}
	fmt.Printf("|%-9s|%-15s|%-15s|%-15s|%-15s|\n", "id", "Name", "Power", "defense", "Speed")
	fmt.Println("___________________________________________________________________________")
	for rows.Next() {
		rows.Scan(&attacks.Id, &attacks.Name, &attacks.Power, &attacks.Defense, &attacks.Speed)
		fmt.Printf("|%-9d|%-15s|%-15d|%-15d|%-15d|\n", attacks.Id, attacks.Name, attacks.Power, attacks.Defense, attacks.Speed)
	}
}

func UpdateAttacks(db *sql.DB) int64 {
	var statement, err = db.Prepare("UPDATE attacks SET name = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	var attacks = Attack{}
	fmt.Println("Enter id")
	fmt.Scan(&attacks.Id)
	fmt.Println("Enter name to update")
	fmt.Scan(&attacks.Name)

	var res, _ = statement.Exec(attacks.Name, attacks.Id)
	var n, _ = res.RowsAffected()

	return n
}

func DeleteAttacks(db *sql.DB) int64 {
	var statement, err = db.Prepare("DELETE from attacks WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	var attacks = Attack{}
	fmt.Println("Enter id")
	fmt.Scan(&attacks.Id)
	var res, _ = statement.Exec(attacks.Id)
	var n, _ = res.RowsAffected()

	return n
}

func SearchAttacks(db *sql.DB, id int) (attacks *Attack, err error) {
	var aux Attack
	var row = db.QueryRow("SELECT id, name, power, defense, speed FROM attacks WHERE id = ?", id)
	err = row.Scan(&aux.Id, &aux.Name, &aux.Power, &aux.Defense, &aux.Speed)
	if err != nil {
		return
	}
	attacks = &aux
	return
}
