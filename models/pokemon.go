package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sandjuarezg/sqlite-pokemon/function"
)

type Pokemon struct {
	Id    int
	Name  string
	Type  string
	Level int
}

func AddPokemon(db *sql.DB) (err error) {
	statement, err := db.Prepare("INSERT INTO pokemons (name, type, level) VALUES (?, ?, ?)")
	if err != nil {
		err = function.ErrInsert
		return
	}
	defer statement.Close()

	var poke = Pokemon{}
	fmt.Println("Enter a name")
	fmt.Scan(&poke.Name)

	var exit bool
	var opc int

	for !exit {
		fmt.Println("Select a type")
		fmt.Println("1. Earth")
		fmt.Println("2. Fire")
		fmt.Println("3. Air")
		fmt.Println("4. Water")
		fmt.Println("5. Normal")
		fmt.Println("6. Electric")
		fmt.Println("7. Plant")
		fmt.Println("8. Legendary")

		fmt.Scan(&opc)

		switch opc {
		case 1:
			poke.Type = "Earth"
			exit = true
		case 2:
			poke.Type = "Fire"
			exit = true
		case 3:
			poke.Type = "Air"
			exit = true
		case 4:
			poke.Type = "Water"
			exit = true
		case 5:
			poke.Type = "Normal"
			exit = true
		case 6:
			poke.Type = "Electric"
			exit = true
		case 7:
			poke.Type = "Plant"
			exit = true
		case 8:
			poke.Type = "Legendary"
			exit = true
		default:
			fmt.Println("Option not valid")
			function.CleanConsole(2)
		}
	}
	fmt.Println("Enter a lever")
	fmt.Scan(&poke.Level)
	statement.Exec(poke.Name, poke.Type, poke.Level)

	return
}

func ShowPokemon(db *sql.DB) (err error) {
	rows, err := db.Query("SELECT id, name, type, level FROM pokemons")
	if err != nil {
		err = function.ErrShowData
		return
	}
	defer rows.Close()

	var poke = Pokemon{}
	fmt.Printf("|%-7s|%-15s|%-15s|%-7s|\n", "id", "Name", "Type", "Level")
	fmt.Println("________________________________________________")
	for rows.Next() {
		err = rows.Scan(&poke.Id, &poke.Name, &poke.Type, &poke.Level)
		if err != nil {
			err = function.ErrScan
			return
		}
		fmt.Printf("|%-7d|%-15s|%-15s|%-7d|\n", poke.Id, poke.Name, poke.Type, poke.Level)
	}
	return
}

func UpdatePokemon(db *sql.DB) (n int64, err error) {
	statement, err := db.Prepare("UPDATE pokemons SET level = ? WHERE id = ?")
	if err != nil {
		err = function.ErrUpdate
		return
	}
	defer statement.Close()

	var poke = Pokemon{}
	fmt.Println("Enter id")
	fmt.Scan(&poke.Id)
	fmt.Println("Enter level to update")
	fmt.Scan(&poke.Level)

	var res, _ = statement.Exec(poke.Level, poke.Id)
	n, _ = res.RowsAffected()

	return
}

func DeletePokemon(db *sql.DB) (n int64, err error) {
	statement, err := db.Prepare("DELETE from pokemons WHERE id = ?")
	if err != nil {
		err = function.ErrDelete
		return
	}
	defer statement.Close()

	var poke = Pokemon{}
	fmt.Println("Enter id")
	fmt.Scan(&poke.Id)
	var res, _ = statement.Exec(poke.Id)
	n, _ = res.RowsAffected()

	return
}

func SearchPokemon(db *sql.DB, id int) (pokemon *Pokemon, err error) {
	var aux Pokemon
	var row = db.QueryRow("SELECT id, name, type, level FROM pokemons WHERE id = ?", id)
	err = row.Scan(&aux.Id, &aux.Name, &aux.Type, &aux.Level)
	if err != nil {
		err = function.ErrScan
		return
	}
	pokemon = &aux
	return
}
