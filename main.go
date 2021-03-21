package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sandjuarezg/sqlite-pokemon/function"
	"github.com/sandjuarezg/sqlite-pokemon/models"
)

func main() {
	function.SqlMigration()

	var db, err = sql.Open("sqlite3", "./pokemon.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var opc int
	var exit bool

	for !exit {
		fmt.Println("-- Welcome to Sand's Pokemon --")
		fmt.Println("0. Exit")
		fmt.Println("1. Add")
		fmt.Println("2. Show")
		fmt.Println("3. Update")
		fmt.Println("4. Delete")
		fmt.Println("5. Go to relation")
		fmt.Scan(&opc)

		switch opc {
		case 0:
			fmt.Println("E X I T . . .")
			exit = true
		case 1:
			function.CleanConsole(1)
			var back bool

			for !back {
				fmt.Println("-- Add --")
				fmt.Println("0. Back")
				fmt.Println("1. User")
				fmt.Println("2. Pokemon")
				fmt.Println("3. Attack")

				fmt.Scan(&opc)

				switch opc {
				case 0:
					fmt.Println("B A C K . . .")
					back = true
					function.CleanConsole(2)
				case 1:
					models.AddUser(db)
					function.CleanConsole(1)
				case 2:
					models.AddPokemon(db)
					function.CleanConsole(1)
				case 3:
					models.AddAttacks(db)
					function.CleanConsole(1)
				default:
					fmt.Println("Option not valid")
					function.CleanConsole(2)
				}
			}
		case 2:
			function.CleanConsole(1)
			var back bool

			for !back {
				fmt.Println("-- Show --")
				fmt.Println("0. Back")
				fmt.Println("1. User")
				fmt.Println("2. Pokemon")
				fmt.Println("3. Attack")
				fmt.Scan(&opc)

				switch opc {
				case 0:
					fmt.Println("B A C K . . .")
					back = true
					function.CleanConsole(2)
				case 1:
					models.ShowUser(db)
					function.CleanConsole(4)
				case 2:
					models.ShowPokemon(db)
					function.CleanConsole(4)
				case 3:
					models.ShowAttacks(db)
					function.CleanConsole(4)
				default:
					fmt.Println("Option not valid")
					function.CleanConsole(2)
				}
			}
		case 3:
			function.CleanConsole(1)
			var back bool

			for !back {
				fmt.Println("-- Update --")
				fmt.Println("0. Back")
				fmt.Println("1. User")
				fmt.Println("2. Pokemon")
				fmt.Println("3. Attack")
				fmt.Scan(&opc)

				switch opc {
				case 0:
					fmt.Println("B A C K . . .")
					back = true
					function.CleanConsole(2)
				case 1:
					var n = models.UpdateUser(db)
					if n == 0 {
						fmt.Println("Problem: Not found id")
					}
					function.CleanConsole(2)
				case 2:
					var n = models.UpdatePokemon(db)
					if n == 0 {
						fmt.Println("Problem: Not found id")
					}
					function.CleanConsole(2)
				case 3:
					var n = models.UpdateAttacks(db)
					if n == 0 {
						fmt.Println("Problem: Not found id")
					}
					function.CleanConsole(2)
				default:
					fmt.Println("Option not valid")
					function.CleanConsole(2)
				}
			}
		case 4:
			function.CleanConsole(1)
			var back bool

			for !back {
				fmt.Println("-- Delete --")
				fmt.Println("0. Back")
				fmt.Println("1. User")
				fmt.Println("2. Pokemon")
				fmt.Println("3. Attack")
				fmt.Scan(&opc)

				switch opc {
				case 0:
					fmt.Println("B A C K . . .")
					back = true
					function.CleanConsole(2)
				case 1:
					var n = models.DeleteUser(db)
					if n == 0 {
						fmt.Println("Problem: Not found id")
					}
					function.CleanConsole(2)
				case 2:
					var n = models.DeletePokemon(db)
					if n == 0 {
						fmt.Println("Problem: Not found id")
					}
					function.CleanConsole(2)
				case 3:
					var n = models.DeleteAttacks(db)
					if n == 0 {
						fmt.Println("Problem: Not found id")
					}
					function.CleanConsole(2)
				default:
					fmt.Println("Option not valid")
					function.CleanConsole(2)
				}
			}
		case 5:
			function.CleanConsole(1)
			var back bool

			for !back {
				fmt.Println("0. Back")
				fmt.Println("1. Users and Pokemons")
				fmt.Println("2. Pokemons and Attacks")
				fmt.Scan(&opc)

				switch opc {
				case 0:
					fmt.Println("B A C K . . .")
					back = true
					function.CleanConsole(2)
				case 1:
					function.CleanConsole(1)
					var back bool

					for !back {
						fmt.Println("0. Back")
						fmt.Println("1. Add pokemon to user")
						fmt.Println("2. Show all users and pokemons")
						fmt.Println("3. Show specific user and his pokemons")
						fmt.Println("4. Delete pokemon to user")
						fmt.Scan(&opc)

						switch opc {
						case 0:
							fmt.Println("B A C K . . .")
							back = true
							function.CleanConsole(2)
						case 1:
							var err = models.AddUserPokemon(db)
							if err != nil {
								fmt.Println("Problem:", err)
							}
							function.CleanConsole(2)
						case 2:
							models.ShowUserPokemonAll(db)
							function.CleanConsole(4)
						case 3:
							var err = models.ShowUserPokemonSpecific(db)
							if err != nil {
								fmt.Println(err)
							}
							function.CleanConsole(4)
						case 4:
							var n = models.DeleteUserPokemon(db)
							if n == 0 {
								fmt.Println("Problem: Not found id")
							}
							function.CleanConsole(2)
						default:
							fmt.Println("Option not valid")
							function.CleanConsole(2)
						}
					}
					function.CleanConsole(1)
				case 2:
					function.CleanConsole(1)
					var back bool

					for !back {
						fmt.Println("0. Back")
						fmt.Println("1. Add attack to pokemon")
						fmt.Println("2. Show all pokemons and attacks")
						fmt.Println("3. Show specific pokemon and his attacks")
						fmt.Println("4. Delete attack to pokemon")
						fmt.Scan(&opc)

						switch opc {
						case 0:
							fmt.Println("B A C K . . .")
							back = true
							function.CleanConsole(2)
						case 1:
							var err = models.AddPokemonAttack(db)
							if err != nil {
								fmt.Println("Problem:", err)
							}
							function.CleanConsole(2)
						case 2:
							models.ShowPokemonAttackAll(db)
							function.CleanConsole(4)
						case 3:
							var err = models.ShowPokemonAttackSpecific(db)
							if err != nil {
								fmt.Println(err)
							}
							function.CleanConsole(4)
						case 4:
							var n = models.DeletePokemonAttack(db)
							if n == 0 {
								fmt.Println("Problem: Not found id")
							}
							function.CleanConsole(2)
						default:
							fmt.Println("Option not valid")
							function.CleanConsole(2)
						}
					}
					function.CleanConsole(1)
				default:
					fmt.Println("Option not valid")
					function.CleanConsole(2)
				}
			}
		default:
			fmt.Println("Option not valid")
			function.CleanConsole(2)
		}
	}
}
