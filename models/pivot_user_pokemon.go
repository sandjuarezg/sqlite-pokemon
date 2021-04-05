package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sandjuarezg/sqlite-pokemon/function"
)

type UserPokemon struct {
	IdUser    int
	IdPokemon int
}

func AddUserPokemon(db *sql.DB) (err error) {
	var user_pokemon = UserPokemon{}
	statement, err := db.Prepare("INSERT INTO user_pokemon (id_user, id_pokemon) VALUES (?, ?)")
	if err != nil {
		err = function.ErrInsert
		return
	}
	defer statement.Close()

	fmt.Println("Enter user id")
	fmt.Scan(&user_pokemon.IdUser)
	_, err = SearchUser(db, user_pokemon.IdUser)
	if err != nil {
		err = function.ErrUnknown
		return
	}

	fmt.Println("Enter pokemon id")
	fmt.Scan(&user_pokemon.IdPokemon)
	_, err = SearchPokemon(db, user_pokemon.IdPokemon)
	if err != nil {
		err = function.ErrUnknown
		return
	}

	//Check not to repeat
	var aux = UserPokemon{}
	var row = db.QueryRow(`
		SELECT 
			users.id 
			FROM 
				user_pokemon 
			INNER JOIN users ON users.id = user_pokemon.id_user 
			WHERE 
				users.id = ? AND user_pokemon.id_pokemon = ? 
		`, user_pokemon.IdUser, user_pokemon.IdPokemon)
	err = row.Scan(&aux.IdUser)
	if err != nil {
		//If no data found, then I can insert
		statement.Exec(user_pokemon.IdUser, user_pokemon.IdPokemon)
		err = nil
		return
	}
	err = function.ErrDuplicate
	return
}

func ShowUserPokemonAll(db *sql.DB) (err error) {
	rows, err := db.Query(`
		SELECT 
			users.id, users.name, users.password, users.ocupation, pokemons.id, pokemons.name, pokemons.type, pokemons.level 
			FROM 
				user_pokemon
			INNER JOIN users ON users.id = user_pokemon.id_user 
			INNER JOIN pokemons ON pokemons.id = user_pokemon.id_pokemon 
			ORDER BY 
				users.id ASC
		`)
	if err != nil {
		err = function.ErrShowData
		return
	}
	defer rows.Close()

	var user = User{}
	var pokemon = Pokemon{}
	fmt.Printf("|%-7s|%-15s|%-10s|%-15s|%-10s|%-10s|%-10s|%-10s|\n", "id_user", "Name", "Password", "Ocupation", "id_pokemon", "Name", "Type", "Level")
	fmt.Println("_________________________________________________________________________________________")
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Pass, &user.Ocupa, &pokemon.Id, &pokemon.Name, &pokemon.Type, &pokemon.Level)
		if err != nil {
			err = function.ErrScan
			return
		}
		fmt.Printf("|%-7d|%-15s|%-10s|%-15s|%-10d|%-10s|%-10s|%-10d|\n", user.Id, user.Name, user.Pass, user.Ocupa, pokemon.Id, pokemon.Name, pokemon.Type, pokemon.Level)
	}
	return
}

func ShowUserPokemonSpecific(db *sql.DB) (err error) {
	var user_pokemon = UserPokemon{}
	fmt.Println("Enter user id")
	fmt.Scan(&user_pokemon.IdUser)

	row, err := db.Query(`
		SELECT 
			users.id, users.name, users.password, users.ocupation, pokemons.id, pokemons.name, pokemons.type, pokemons.level 
			FROM 
				user_pokemon
			INNER JOIN users ON users.id = user_pokemon.id_user 
			INNER JOIN pokemons ON pokemons.id = user_pokemon.id_pokemon 
			WHERE 
				user_pokemon.id_user = ? 
		`, user_pokemon.IdUser)
	if err != nil {
		err = function.ErrShowData
		return
	}
	defer row.Close()

	var user = User{}
	var pokemon = Pokemon{}
	fmt.Printf("|%-7s|%-15s|%-10s|%-15s|%-10s|%-10s|%-10s|%-10s|\n", "id_user", "Name", "Password", "Ocupation", "id_pokemon", "Name", "Type", "Level")
	fmt.Println("_________________________________________________________________________________________")
	for row.Next() {
		err = row.Scan(&user.Id, &user.Name, &user.Pass, &user.Ocupa, &pokemon.Id, &pokemon.Name, &pokemon.Type, &pokemon.Level)
		if err != nil {
			err = function.ErrScan
			return
		}
		fmt.Printf("|%-7d|%-15s|%-10s|%-15s|%-10d|%-10s|%-10s|%-10d|\n", user.Id, user.Name, user.Pass, user.Ocupa, pokemon.Id, pokemon.Name, pokemon.Type, pokemon.Level)
	}
	return
}

func DeleteUserPokemon(db *sql.DB) (n int64, err error) {
	statement, err := db.Prepare("DELETE from user_pokemon WHERE id_user = ? AND id_pokemon = ?")
	if err != nil {
		err = function.ErrDelete
		return
	}
	defer statement.Close()

	var user_pokemon = UserPokemon{}
	fmt.Println("Enter user id")
	fmt.Scan(&user_pokemon.IdUser)
	fmt.Println("Enter pokemon id")
	fmt.Scan(&user_pokemon.IdPokemon)

	var res, _ = statement.Exec(user_pokemon.IdUser, user_pokemon.IdPokemon)
	n, _ = res.RowsAffected()

	return
}
