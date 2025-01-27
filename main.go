package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/CrymsonShadows/blog-aggregator/internal/config"
	"github.com/CrymsonShadows/blog-aggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading from config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	dbQueries := database.New(db)
	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: map[string]func(*state, command) error{},
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", midllewareLoggedIn(hanlderAddFeed))
	cmds.register("feeds", handlerListFeeds)
	cmds.register("follow", midllewareLoggedIn(handlerFollow))
	cmds.register("following", midllewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", midllewareLoggedIn(handlerUnfollow))

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
