package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/SplinterSword/GOGator/internal/config"
	"github.com/SplinterSword/GOGator/internal/database"
	_ "github.com/lib/pq"
)

type State struct {
	db            *database.Queries
	CurrentConfig *config.Config
}

func main() {
	cfg := State{}
	config, err := config.Read()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := sql.Open("postgres", config.DBURL)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	cfg.db = database.New(db)
	cfg.CurrentConfig = &config

	commands := Commands{
		CommandMap: make(map[string]CommandHandler),
	}
	commands.Register("login", handleLogin)
	commands.Register("register", handleRegister)
	commands.Register("reset", handleReset)
	commands.Register("users", handleGetUsers)
	commands.Register("agg", handleAgg)
	commands.Register("addfeed", handleAddFeeds)
	commands.Register("feeds", handleFeeds)

	Args := os.Args

	if len(Args) < 2 && (Args[1] != "reset" && Args[1] != "getUsers" && Args[1] != "feeds") {
		log.Fatal("Invalid number of arguments")
		return
	}

	command := Command{
		Name: Args[1],
		Args: Args[2:],
	}

	err = commands.Run(&cfg, command)
	if err != nil {
		log.Fatal(err.Error())
	}
}
