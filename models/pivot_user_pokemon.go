package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User_pokemon struct {
	Id_user    int
	Id_pokemon int
}

func AddUserPokemon(db *sql.DB) error {
	var user_pokemon = User_pokemon{}

	fmt.Println("Enter user id")
	fmt.Scan(&user_pokemon.Id_user)
	var _, err = SearchUser(db, user_pokemon.Id_user)
	if err != nil {
		return errors.New("No user found")
	}

	fmt.Println("Enter pokemon id")
	fmt.Scan(&user_pokemon.Id_pokemon)
	_, err = SearchPokemon(db, user_pokemon.Id_pokemon)
	if err != nil {
		return errors.New("No pokemon found")
	}

	//Check not to repeat
	var aux = User_pokemon{}
	var row = db.QueryRow("SELECT id_user, id_pokemon FROM user_pokemon WHERE id_user = ? AND id_pokemon = ?", user_pokemon.Id_user, user_pokemon.Id_pokemon)
	err = row.Scan(&aux.Id_user, &aux.Id_pokemon)
	if err != nil {
		//If no data found, then I can insert
		statement, err := db.Prepare("INSERT INTO user_pokemon (id_user, id_pokemon) VALUES (?, ?)")
		if err != nil {
			log.Fatal(err)
		}
		defer statement.Close()
		statement.Exec(user_pokemon.Id_user, user_pokemon.Id_pokemon)
		return nil
	}
	return errors.New("The user already has this pokemon")
}

func ShowUserPokemonAll(db *sql.DB) {
	var rows, err = db.Query("SELECT id_user, id_pokemon FROM user_pokemon ORDER BY id_user ASC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var user_pokemon = User_pokemon{}
	fmt.Printf("|%-7s|%-10s|\n", "id_user", "id_pokemon")
	fmt.Println("____________________")
	for rows.Next() {
		rows.Scan(&user_pokemon.Id_user, &user_pokemon.Id_pokemon)
		fmt.Printf("|%-7d|%-10d|\n", user_pokemon.Id_user, user_pokemon.Id_pokemon)
	}
}

func ShowUserPokemonSpecific(db *sql.DB) error {
	var user_pokemon = User_pokemon{}
	fmt.Println("Enter user id")
	fmt.Scan(&user_pokemon.Id_user)

	var row, err = db.Query("SELECT id_pokemon FROM user_pokemon WHERE id_user = ?", user_pokemon.Id_user)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var user, _ = SearchUser(db, user_pokemon.Id_user)

	fmt.Printf("ID: %d, Name: %s, Password: %s, Ocupation: %s\n", user.Id, user.Name, user.Pass, user.Ocupa)
	fmt.Printf("|%-10s|%-10s|%-10s|%-10s|\n", "id_pokemon", "Name", "Type", "Level")
	fmt.Println("_____________________________________________")
	for row.Next() {
		err = row.Scan(&user_pokemon.Id_pokemon)
		if err != nil {
			return errors.New("User without pokemons")
		}
		pokemon, _ := SearchPokemon(db, user_pokemon.Id_pokemon)
		fmt.Printf("|%-10d|%-10s|%-10s|%-10d|\n", pokemon.Id, pokemon.Name, pokemon.Type, pokemon.Level)
	}
	return nil
}

func DeleteUserPokemon(db *sql.DB) int64 {
	var statement, err = db.Prepare("DELETE from user_pokemon WHERE id_user = ? AND id_pokemon = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	var user_pokemon = User_pokemon{}
	fmt.Println("Enter user id")
	fmt.Scan(&user_pokemon.Id_user)
	fmt.Println("Enter pokemon id")
	fmt.Scan(&user_pokemon.Id_pokemon)

	var res, _ = statement.Exec(user_pokemon.Id_user, user_pokemon.Id_user)
	var n, _ = res.RowsAffected()

	return n
}
