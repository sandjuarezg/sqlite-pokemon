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

func CleanConsole(d time.Duration) {
	fmt.Println("Wait a second . . .")

	time.Sleep(d * time.Second)
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
}
