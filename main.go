package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sainikmandal/gator/internal/config"
	"github.com/sainikmandal/gator/internal/database"
	"os"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("failed to read config:", err)
		os.Exit(1)
	}
	dbURL := cfg.DBURL
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	s := &state{db: dbQueries, cfg: &cfg}

	cmds := &commands{registeredCommands: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAgg)

	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	cmd := command{Name: os.Args[1], Args: os.Args[2:]}
	err = cmds.run(s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
