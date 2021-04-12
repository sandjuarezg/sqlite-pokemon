package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sandjuarezg/sqlite-pokemon/function"
)

type PokemonAttack struct {
	IdPokemon int
	IdAttack  int
}

func AddPokemonAttack(db *sql.DB) (err error) {
	smt, err := db.Prepare("INSERT INTO pokemon_attack (id_pokemon, id_attack) VALUES (?, ?)")
	if err != nil {
		err = function.ErrInsert
		return
	}
	defer smt.Close()

	var pokemonAttack = PokemonAttack{}
	fmt.Println("Enter pokemon id")
	fmt.Scan(&pokemonAttack.IdPokemon)
	_, err = SearchPokemon(db, pokemonAttack.IdPokemon)
	if err != nil {
		err = function.ErrUnknown
		return
	}

	fmt.Println("Enter attack id")
	fmt.Scan(&pokemonAttack.IdAttack)
	_, err = SearchAttacks(db, pokemonAttack.IdAttack)
	if err != nil {
		err = function.ErrUnknown
		return
	}

	//Check not to repeat
	var aux = PokemonAttack{}
	var row = db.QueryRow(`
		SELECT 
			pokemons.id 
			FROM 
				pokemon_attack 
			INNER JOIN pokemons ON pokemon_attack.id_pokemon = pokemons.id
			WHERE 
				pokemons.id = ? AND pokemon_attack.id_attack = ?
		`, pokemonAttack.IdPokemon, pokemonAttack.IdAttack)
	err = row.Scan(&aux.IdPokemon)
	if err != nil {
		//If no data found, then I can insert
		smt.Exec(pokemonAttack.IdPokemon, pokemonAttack.IdAttack)
		err = nil
		return
	}
	err = function.ErrDuplicate
	return
}

func ShowPokemonAttackAll(db *sql.DB) (err error) {
	rows, err := db.Query(`
		SELECT 
			pokemons.id, pokemons.name, pokemons.type, pokemons.level, attacks.id, attacks.name, attacks.power, attacks.defense, attacks.speed 
			FROM 
				pokemon_attack
			INNER JOIN pokemons ON pokemons.id = pokemon_attack.id_pokemon 
			INNER JOIN attacks ON attacks.id = pokemon_attack.id_attack 
			ORDER BY 
				pokemons.id ASC
	`)
	if err != nil {
		err = function.ErrShowData
		return
	}
	defer rows.Close()

	var attack = Attack{}
	var pokemon = Pokemon{}
	fmt.Printf("|%-10s|%-10s|%-10s|%-10s|%-10s|%-12s|%-10s|%-10s|%-10s|\n", "id_pokemon", "Name", "Type", "Level", "id_attack", "Name", "Power", "Defense", "Seep")
	fmt.Println("______________________________________________________________________________________________________")
	for rows.Next() {
		err = rows.Scan(&pokemon.Id, &pokemon.Name, &pokemon.Type, &pokemon.Level, &attack.Id, &attack.Name, &attack.Power, &attack.Defense, &attack.Speed)
		if err != nil {
			err = function.ErrScan
			return
		}
		fmt.Printf("|%-10d|%-10s|%-10s|%-10d|%-10d|%-12s|%-10d|%-10d|%-10d|\n", pokemon.Id, pokemon.Name, pokemon.Type, pokemon.Level, attack.Id, attack.Name, attack.Power, attack.Defense, attack.Speed)
	}
	return
}

func ShowPokemonAttackSpecific(db *sql.DB) (err error) {
	var pokemonAttack = PokemonAttack{}
	fmt.Println("Enter pokemon id")
	fmt.Scan(&pokemonAttack.IdPokemon)

	row, err := db.Query(`
		SELECT 
			pokemons.id, pokemons.name, pokemons.type, pokemons.level, attacks.id, attacks.name, attacks.power, attacks.defense, attacks.speed 
			FROM 
				pokemon_attack
			INNER JOIN pokemons ON pokemons.id = pokemon_attack.id_pokemon 
			INNER JOIN attacks ON attacks.id = pokemon_attack.id_attack
			WHERE 
				pokemon_attack.id_pokemon = ?
		`, pokemonAttack.IdPokemon)
	if err != nil {
		err = function.ErrShowData
		return
	}
	defer row.Close()

	var pokemon = Pokemon{}
	var attack = Attack{}
	fmt.Printf("|%-10s|%-10s|%-10s|%-10s|%-10s|%-12s|%-10s|%-10s|%-10s|\n", "id_pokemon", "Name", "Type", "Level", "id_attack", "Name", "Power", "Defense", "Seep")
	fmt.Println("______________________________________________________________________________________________________")
	for row.Next() {
		err = row.Scan(&pokemon.Id, &pokemon.Name, &pokemon.Type, &pokemon.Level, &attack.Id, &attack.Name, &attack.Power, &attack.Defense, &attack.Speed)
		if err != nil {
			err = function.ErrScan
			return
		}
		fmt.Printf("|%-10d|%-10s|%-10s|%-10d|%-10d|%-12s|%-10d|%-10d|%-10d|\n", pokemon.Id, pokemon.Name, pokemon.Type, pokemon.Level, attack.Id, attack.Name, attack.Power, attack.Defense, attack.Speed)
	}
	return
}

func DeletePokemonAttack(db *sql.DB) (n int64, err error) {
	var pokemonAttack = PokemonAttack{}
	fmt.Println("Enter pokemon id")
	fmt.Scan(&pokemonAttack.IdPokemon)
	fmt.Println("Enter attack id")
	fmt.Scan(&pokemonAttack.IdAttack)

	row, err := db.Exec("DELETE from pokemon_attack WHERE id_pokemon = ? AND id_attack = ?", pokemonAttack.IdPokemon, pokemonAttack.IdAttack)
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
