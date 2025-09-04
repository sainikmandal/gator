package main

import (
	"fmt"
	"github.com/sainikmandal/gator/internal/config"
	"os"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("failed to read config:", err)
		os.Exit(1)
	}
	s := &state{cfg: &cfg}

	cmds := &commands{registeredCommands: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)

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
