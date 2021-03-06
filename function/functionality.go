package function

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var (
	ErrInsert    = errors.New("Problem: error to insert")
	ErrShowData  = errors.New("Problem: error to show data")
	ErrUpdate    = errors.New("Problem: error to update")
	ErrDelete    = errors.New("Problem: error to delete")
	ErrScan      = errors.New("Problem: error to scan rows")
	ErrUnknown   = errors.New("Problem: Unknown data")
	ErrDuplicate = errors.New("Problem: Duplicate data")
)

func CleanConsole(second int) {
	fmt.Println("Wait a second . . .")

	time.Sleep(time.Duration(second) * time.Second)
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
