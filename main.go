package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/saniak-hub/gator/internal/config"
	"github.com/saniak-hub/gator/internal/database"
)

func main() {
	conf := config.Read()

	db, err := sql.Open("postgres", conf.DbURL)
	if err != nil {
		fmt.Println("Error connectind db", err)
	}
	dbQueries := database.New(db)

	newConf := state{
		config: &conf,
		db:     dbQueries,
	}

	cmds := commands{
		cmds: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("browse", middlewareLoggedIn(handlerBrowse))

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Not enough arguments were provided")
		os.Exit(1)
	}

	cmd := command{}
	if len(args) > 2 {
		cmd.name = args[1]
		cmd.args = args[2:]
	} else {
		cmd.name = args[1]
	}
	if err := cmds.run(&newConf, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//fmt.Println(conf)
}
