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
		fmt.Println("1. Add")
		fmt.Println("2. Show")
		fmt.Println("3. Update")
		fmt.Println("4. Delete")
		fmt.Println("5. Go to relation")
		fmt.Println("6. Exit")
		fmt.Scan(&opc)

		switch opc {
		case 1:
			function.CleanConsole()
			var back bool

			for !back {
				submenu("Add")
				fmt.Scan(&opc)

				switch opc {
				case 1:
					models.AddUser(db)
					function.CleanConsole()
				case 2:
					models.AddPokemon(db)
					function.CleanConsole()
				case 3:
					models.AddAttacks(db)
					function.CleanConsole()
				case 4:
					fmt.Println("B A C K . . .")
					back = true
					function.CleanConsole()
				default:
					fmt.Println("Option not valid")
					function.CleanConsole()
				}
			}
		case 2:
			function.CleanConsole()
			var back bool

			for !back {
				submenu("Show")
				fmt.Scan(&opc)

				switch opc {
				case 1:
					models.ShowUser(db)
					function.CleanConsole()
				case 2:
					models.ShowPokemon(db)
					function.CleanConsole()
				case 3:
					models.ShowAttacks(db)
					function.CleanConsole()
				case 4:
					fmt.Println("B A C K . . .")
					back = true
					function.CleanConsole()
				default:
					fmt.Println("Option not valid")
					function.CleanConsole()
				}
			}
		case 3:
			function.CleanConsole()
			var back bool

			for !back {
				submenu("Update")
				fmt.Scan(&opc)

				switch opc {
				case 1:
					var n = models.UpdateUser(db)
					if n == 0 {
						fmt.Println("Problem: Not found id")
					}
					function.CleanConsole()
				case 2:
					var n = models.UpdatePokemon(db)
					if n == 0 {
						fmt.Println("Problem: Not found id")
					}
					function.CleanConsole()
				case 3:
					var n = models.UpdateAttacks(db)
					if n == 0 {
						fmt.Println("Problem: Not found id")
					}
					function.CleanConsole()
				case 4:
					fmt.Println("B A C K . . .")
					back = true
					function.CleanConsole()
				default:
					fmt.Println("Option not valid")
					function.CleanConsole()
				}
			}
		case 4:
			function.CleanConsole()
			var back bool

			for !back {
				submenu("Delete")
				fmt.Scan(&opc)

				switch opc {
				case 1:
					var n = models.DeleteUser(db)
					if n == 0 {
						fmt.Println("Problem: Not found id")
					}
					function.CleanConsole()
				case 2:
					var n = models.DeletePokemon(db)
					if n == 0 {
						fmt.Println("Problem: Not found id")
					}
					function.CleanConsole()
				case 3:
					var n = models.DeleteAttacks(db)
					if n == 0 {
						fmt.Println("Problem: Not found id")
					}
					function.CleanConsole()
				case 4:
					fmt.Println("B A C K . . .")
					back = true
					function.CleanConsole()
				default:
					fmt.Println("Option not valid")
					function.CleanConsole()
				}
			}
		case 5:
			function.CleanConsole()
			var back bool

			for !back {
				fmt.Println("1. Users and Pokemons")
				fmt.Println("2. Pokemons and Attacks")
				fmt.Println("3. Back")
				fmt.Scan(&opc)

				switch opc {
				case 1:
					function.CleanConsole()
					var back bool

					for !back {
						fmt.Println("1. Add pokemon to user")
						fmt.Println("2. Show all users and pokemons")
						fmt.Println("3. Show specific user and his pokemons")
						fmt.Println("4. Delete pokemon to user")
						fmt.Println("5. Back")
						fmt.Scan(&opc)

						switch opc {
						case 1:
							var err = models.AddUserPokemon(db)
							if err != nil {
								fmt.Println("Problem:", err)
							}
							function.CleanConsole()
						case 2:
							models.ShowUserPokemonAll(db)
							function.CleanConsole()
						case 3:
							var err = models.ShowUserPokemonSpecific(db)
							if err != nil {
								fmt.Println(err)
							}
							function.CleanConsole()
						case 4:
							var n = models.DeleteUserPokemon(db)
							if n == 0 {
								fmt.Println("Problem: Not found id")
							}
							function.CleanConsole()
						case 5:
							fmt.Println("B A C K . . .")
							back = true
							function.CleanConsole()
						default:
							fmt.Println("Option not valid")
							function.CleanConsole()
						}
					}
					function.CleanConsole()
				case 2:
					function.CleanConsole()
					var back bool

					for !back {
						fmt.Println("1. Add attack to pokemon")
						fmt.Println("2. Show all pokemons and attacks")
						fmt.Println("3. Show specific pokemon and his attacks")
						fmt.Println("4. Delete attack to pokemon")
						fmt.Println("5. Back")
						fmt.Scan(&opc)

						switch opc {
						case 1:
							var err = models.AddPokemonAttack(db)
							if err != nil {
								fmt.Println("Problem:", err)
							}
							function.CleanConsole()
						case 2:
							models.ShowPokemonAttackAll(db)
							function.CleanConsole()
						case 3:
							var err = models.ShowPokemonAttackSpecific(db)
							if err != nil {
								fmt.Println(err)
							}
							function.CleanConsole()
						case 4:
							var n = models.DeletePokemonAttack(db)
							if n == 0 {
								fmt.Println("Problem: Not found id")
							}
							function.CleanConsole()
						case 5:
							fmt.Println("B A C K . . .")
							back = true
							function.CleanConsole()
						default:
							fmt.Println("Option not valid")
							function.CleanConsole()
						}
					}
					function.CleanConsole()
				case 3:
					fmt.Println("B A C K . . .")
					back = true
					function.CleanConsole()
				default:
					fmt.Println("Option not valid")
					function.CleanConsole()
				}
			}
		case 6:
			fmt.Println("E X I T . . .")
			exit = true
		default:
			fmt.Println("Option not valid")
			function.CleanConsole()
		}
	}
}

func submenu(str string) {
	fmt.Printf("-- %s --\n", str)
	fmt.Println("1. User")
	fmt.Println("2. Pokemon")
	fmt.Println("3. Attack")
	fmt.Println("4. Back")
}
