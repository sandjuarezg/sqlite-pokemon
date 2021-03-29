package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type PokemonAttack struct {
	Id_pokemon int
	Id_attack  int
}

func AddPokemonAttack(db *sql.DB) (err error) {
	var pokemon_attack = PokemonAttack{}

	fmt.Println("Enter pokemon id")
	fmt.Scan(&pokemon_attack.Id_pokemon)
	_, err = SearchPokemon(db, pokemon_attack.Id_pokemon)
	if err != nil {
		err = errors.New("No pokemon found")
		return
	}

	fmt.Println("Enter attack id")
	fmt.Scan(&pokemon_attack.Id_attack)
	_, err = SearchAttack(db, pokemon_attack.Id_attack)
	if err != nil {
		err = errors.New("No attack found")
		return
	}

	//Check not to repeat
	var aux = PokemonAttack{}
	var row = db.QueryRow("SELECT id_pokemon, id_attack FROM pokemon_attack WHERE id_pokemon = ? AND id_attack = ?", pokemon_attack.Id_pokemon, pokemon_attack.Id_attack)
	err = row.Scan(&aux.Id_pokemon, &aux.Id_attack)
	if err != nil {
		//If no data found, then I can insert
		statement, err := db.Prepare("INSERT INTO pokemon_attack (id_pokemon, id_attack) VALUES (?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer statement.Close()
		statement.Exec(pokemon_attack.Id_pokemon, pokemon_attack.Id_attack)
		return nil
	}
	err = errors.New("The pokemon already has this attack")
	return
}

func ShowPokemonAttackAll(db *sql.DB) {
	var rows, err = db.Query("SELECT id_pokemon, id_attack FROM pokemon_attack ORDER BY id_pokemon ASC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var pokemon_attack = PokemonAttack{}
	fmt.Printf("|%-10s|%-10s|%-10s|%-10s|%-10s|%-12s|%-10s|%-10s|%-10s|\n", "id_pokemon", "Name", "Type", "Level", "id_attack", "Name", "Power", "Defense", "Seep")
	fmt.Println("______________________________________________________________________________________________________")
	for rows.Next() {
		rows.Scan(&pokemon_attack.Id_pokemon, &pokemon_attack.Id_attack)
		if err != nil {
			log.Fatal(err)
		}
		pokemon, err := SearchPokemon(db, pokemon_attack.Id_pokemon)
		if err != nil {
			log.Fatal(err)
		}
		attack, err := SearchAttack(db, pokemon_attack.Id_attack)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("|%-10d|%-10s|%-10s|%-10d|%-10d|%-12s|%-10d|%-10d|%-10d|\n", pokemon.Id, pokemon.Name, pokemon.Type, pokemon.Level, attack.Id, attack.Name, attack.Power, attack.Defense, attack.Speed)
	}
}

func ShowPokemonAttackSpecific(db *sql.DB) (err error) {
	var pokemon_attack = PokemonAttack{}
	fmt.Println("Enter pokemon id")
	fmt.Scan(&pokemon_attack.Id_pokemon)

	row, err := db.Query("SELECT id_attack FROM pokemon_attack WHERE id_pokemon = ?", pokemon_attack.Id_pokemon)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	pokemon, err := SearchPokemon(db, pokemon_attack.Id_pokemon)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID: %d, Name: %s, Type: %s, Level: %d\n", pokemon.Id, pokemon.Name, pokemon.Type, pokemon.Level)
	fmt.Printf("|%-10s|%-10s|%-10s|%-10s|%-10s|\n", "id_attack", "Name", "Power", "Defense", "Speed")
	fmt.Println("_______________________________________________________")
	for row.Next() {
		err = row.Scan(&pokemon_attack.Id_attack)
		if err != nil {
			err = errors.New("Pokemon without attacks")
			return
		}
		attack, err := SearchAttack(db, pokemon_attack.Id_attack)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("|%-10d|%-10s|%-10d|%-10d|%-10d|\n", attack.Id, attack.Name, attack.Power, attack.Defense, attack.Speed)
	}
	return
}

func DeletePokemonAttack(db *sql.DB) (n int64) {
	var statement, err = db.Prepare("DELETE from pokemon_attack WHERE id_pokemon = ? AND id_attack = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	var pokemon_attack = PokemonAttack{}
	fmt.Println("Enter pokemon id")
	fmt.Scan(&pokemon_attack.Id_pokemon)
	fmt.Println("Enter attack id")
	fmt.Scan(&pokemon_attack.Id_attack)

	var res, _ = statement.Exec(pokemon_attack.Id_pokemon, pokemon_attack.Id_attack)
	n, _ = res.RowsAffected()

	return
}
