package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type UserPokemon struct {
	Id_user    int
	Id_pokemon int
}

func AddUserPokemon(db *sql.DB) (err error) {
	var user_pokemon = UserPokemon{}

	fmt.Println("Enter user id")
	fmt.Scan(&user_pokemon.Id_user)
	_, err = SearchUser(db, user_pokemon.Id_user)
	if err != nil {
		err = errors.New("No user found")
		return
	}

	fmt.Println("Enter pokemon id")
	fmt.Scan(&user_pokemon.Id_pokemon)
	_, err = SearchPokemon(db, user_pokemon.Id_pokemon)
	if err != nil {
		err = errors.New("No pokemon found")
		return
	}

	//Check not to repeat
	var aux = UserPokemon{}
	var row = db.QueryRow("SELECT id FROM users INNER JOIN user_pokemon ON users.id = ? AND user_pokemon.id_pokemon = ?", user_pokemon.Id_user, user_pokemon.Id_pokemon)
	err = row.Scan(&aux.Id_user)
	if err != nil {
		//If no data found, then I can insert
		var statement, err = db.Prepare("INSERT INTO user_pokemon (id_user, id_pokemon) VALUES (?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer statement.Close()
		statement.Exec(user_pokemon.Id_user, user_pokemon.Id_pokemon)
		return nil
	}
	err = errors.New("The user already has this pokemon")
	return
}

func ShowUserPokemonAll(db *sql.DB) {
	var rows, err = db.Query("SELECT users.id, users.name, users.password, users.ocupation, pokemons.id, pokemons.name, pokemons.type, pokemons.level FROM users, pokemons INNER JOIN user_pokemon WHERE users.id = user_pokemon.id_user AND pokemons.id = user_pokemon.id_pokemon ORDER BY users.id ASC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var user = User{}
	var pokemon = Pokemon{}
	fmt.Printf("|%-7s|%-15s|%-10s|%-15s|%-10s|%-10s|%-10s|%-10s|\n", "id_user", "Name", "Password", "Ocupation", "id_pokemon", "Name", "Type", "Level")
	fmt.Println("_________________________________________________________________________________________")
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Pass, &user.Ocupa, &pokemon.Id, &pokemon.Name, &pokemon.Type, &pokemon.Level)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("|%-7d|%-15s|%-10s|%-15s|%-10d|%-10s|%-10s|%-10d|\n", user.Id, user.Name, user.Pass, user.Ocupa, pokemon.Id, pokemon.Name, pokemon.Type, pokemon.Level)
	}
}

func ShowUserPokemonSpecific(db *sql.DB) (err error) {
	var user_pokemon = UserPokemon{}
	fmt.Println("Enter user id")
	fmt.Scan(&user_pokemon.Id_user)

	row, err := db.Query("SELECT users.id, users.name, users.password, users.ocupation, pokemons.id, pokemons.name, pokemons.type, pokemons.level FROM users, pokemons INNER JOIN user_pokemon WHERE user_pokemon.id_user = ? AND users.id = user_pokemon.id_user AND pokemons.id = user_pokemon.id_pokemon", user_pokemon.Id_user)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	var user = User{}
	var pokemon = Pokemon{}
	fmt.Printf("|%-7s|%-15s|%-10s|%-15s|%-10s|%-10s|%-10s|%-10s|\n", "id_user", "Name", "Password", "Ocupation", "id_pokemon", "Name", "Type", "Level")
	fmt.Println("_________________________________________________________________________________________")
	for row.Next() {
		err = row.Scan(&user.Id, &user.Name, &user.Pass, &user.Ocupa, &pokemon.Id, &pokemon.Name, &pokemon.Type, &pokemon.Level)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("|%-7d|%-15s|%-10s|%-15s|%-10d|%-10s|%-10s|%-10d|\n", user.Id, user.Name, user.Pass, user.Ocupa, pokemon.Id, pokemon.Name, pokemon.Type, pokemon.Level)
	}
	return
}

func DeleteUserPokemon(db *sql.DB) (n int64) {
	var statement, err = db.Prepare("DELETE from user_pokemon WHERE id_user = ? AND id_pokemon = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	var user_pokemon = UserPokemon{}
	fmt.Println("Enter user id")
	fmt.Scan(&user_pokemon.Id_user)
	fmt.Println("Enter pokemon id")
	fmt.Scan(&user_pokemon.Id_pokemon)

	var res, _ = statement.Exec(user_pokemon.Id_user, user_pokemon.Id_pokemon)
	n, _ = res.RowsAffected()

	return
}
