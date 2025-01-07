package main

import (
	"fmt"
	"os"

	"github.com/CrymsonShadows/blog-aggregator/internal/config"
)

func main() {
	cfg := config.Read()
	s := &state{
		cfg: &cfg,
	}
	cmds := commands{
		commandsMap: map[string]func(*state, command) error{},
	}
	cmds.register("login", handlerLogin)
	cmdLineArgs := os.Args
	if len(cmdLineArgs) < 3 {
		fmt.Println("need more arguments")
		os.Exit(1)
	}
	cmdName := cmdLineArgs[1]
	cmdArgs := cmdLineArgs[2:]
	cmd := command{
		name: cmdName,
		args: cmdArgs,
	}
	err := cmds.run(s, cmd)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
