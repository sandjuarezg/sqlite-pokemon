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
	smt, err := db.Prepare("INSERT INTO pokemons (name, type, level) VALUES (?, ?, ?)")
	if err != nil {
		err = function.ErrInsert
		return
	}
	defer smt.Close()

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
	_, err = smt.Exec(poke.Name, poke.Type, poke.Level)
	if err != nil {
		err = function.ErrInsert
		return
	}

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
	var poke = Pokemon{}
	fmt.Println("Enter id")
	fmt.Scan(&poke.Id)
	fmt.Println("Enter level to update")
	fmt.Scan(&poke.Level)

	row, err := db.Exec("UPDATE pokemons SET level = ? WHERE id = ?", poke.Level, poke.Id)
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

func DeletePokemon(db *sql.DB) (n int64, err error) {
	var poke = Pokemon{}
	fmt.Println("Enter id")
	fmt.Scan(&poke.Id)

	row, err := db.Exec("DELETE from pokemons WHERE id = ?", poke.Id)
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

func SearchPokemon(db *sql.DB, id int) (pokemon *Pokemon, err error) {
	pokemon = new(Pokemon)
	var row = db.QueryRow("SELECT id FROM pokemons WHERE id = ?", id)
	err = row.Scan(&pokemon.Id)
	if err != nil {
		err = function.ErrScan
		return
	}
	return
}
