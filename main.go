package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CrymsonShadows/blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading from config: %v", err)
	}
	programState := &state{
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: map[string]func(*state, command) error{},
	}
	cmds.register("login", handlerLogin)

	cmdLineArgs := os.Args
	if len(cmdLineArgs) < 2 {
		fmt.Println("need more arguments")
		os.Exit(1)
	}
	cmdName := cmdLineArgs[1]
	cmdArgs := cmdLineArgs[2:]
	cmd := command{
		name: cmdName,
		args: cmdArgs,
	}
	err = cmds.run(programState, cmd)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
