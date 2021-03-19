package function

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func CleanConsole() {
	fmt.Println("Wait a second . . .")

	time.Sleep(4 * time.Second)
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func SqlMigration() {
	//Check migraton.sql
	var _, err = os.Stat("./migration.sql")
	if os.IsNotExist(err) {
		log.Fatal(err)
	}

	//Get content
	file, _ := os.Open("./migration.sql")
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//Check pokemon.db
	_, err = os.Stat("./pokemon.db")
	if os.IsNotExist(err) {
		var file, err = os.Create("./pokemon.db")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

	}

	//Check table
	db, err := sql.Open("sqlite3", "./pokemon.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Query("SELECT * from users")
	if err != nil {
		_, err = db.Exec(string(content))
		if err != nil {
			log.Fatal(err)
		}
	}
	/*

		SELECT count(*) AS TOTALNUMBEROFTABLES FROM sqlite_master WHERE type = 'table'

				//Check all tables (5)
				if rows.Next() {
					rows.Scan(&res)
				}
				if res != 5 {
					fmt.Println("entra 2")
					_, err = db.Exec(string(content))
					if err != nil {
						log.Fatal(err)
					}
				}
	*/
}
