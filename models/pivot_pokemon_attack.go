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
	var row = db.QueryRow("SELECT id FROM pokemons INNER JOIN pokemon_attack ON pokemons.id = ? AND pokemon_attack.id_attack = ?", pokemon_attack.Id_pokemon, pokemon_attack.Id_attack)
	err = row.Scan(&aux.Id_pokemon)
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
	var rows, err = db.Query("SELECT pokemons.id, pokemons.name, pokemons.type, pokemons.level, attacks.id, attacks.name, attacks.power, attacks.defense, attacks.speed FROM pokemons, attacks INNER JOIN pokemon_attack WHERE pokemons.id = pokemon_attack.id_pokemon AND attacks.id = pokemon_attack.id_attack ORDER BY pokemons.id ASC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var attack = Attack{}
	var pokemon = Pokemon{}
	fmt.Printf("|%-10s|%-10s|%-10s|%-10s|%-10s|%-12s|%-10s|%-10s|%-10s|\n", "id_pokemon", "Name", "Type", "Level", "id_attack", "Name", "Power", "Defense", "Seep")
	fmt.Println("______________________________________________________________________________________________________")
	for rows.Next() {
		err = rows.Scan(&pokemon.Id, &pokemon.Name, &pokemon.Type, &pokemon.Level, &attack.Id, &attack.Name, &attack.Power, &attack.Defense, &attack.Speed)
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

	row, err := db.Query("SELECT pokemons.id, pokemons.name, pokemons.type, pokemons.level, attacks.id, attacks.name, attacks.power, attacks.defense, attacks.speed FROM pokemons, attacks INNER JOIN pokemon_attack WHERE pokemon_attack.id_pokemon = ? AND pokemons.id = pokemon_attack.id_pokemon AND attacks.id = pokemon_attack.id_attack", pokemon_attack.Id_pokemon)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	var pokemon = Pokemon{}
	var attack = Attack{}
	fmt.Printf("|%-10s|%-10s|%-10s|%-10s|%-10s|%-12s|%-10s|%-10s|%-10s|\n", "id_pokemon", "Name", "Type", "Level", "id_attack", "Name", "Power", "Defense", "Seep")
	fmt.Println("______________________________________________________________________________________________________")
	for row.Next() {
		err = row.Scan(&pokemon.Id, &pokemon.Name, &pokemon.Type, &pokemon.Level, &attack.Id, &attack.Name, &attack.Power, &attack.Defense, &attack.Speed)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("|%-10d|%-10s|%-10s|%-10d|%-10d|%-12s|%-10d|%-10d|%-10d|\n", pokemon.Id, pokemon.Name, pokemon.Type, pokemon.Level, attack.Id, attack.Name, attack.Power, attack.Defense, attack.Speed)
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
