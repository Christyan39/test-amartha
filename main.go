package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"test/amartha/api"
	"test/amartha/config"
	"test/amartha/database"
	"test/amartha/usecase"
)

func main() {
	os.Exit(Main())
}

//Main is the main function
func Main() int {

	c := config.Config{
		Server: &config.Server{
			Port: "10001",
		},
		Database: &config.Database{
			Credential: "tyan:@tcp(127.0.0.1:3306)/christyan?parseTime=true",
		},
	}

	db, err := database.Init(*c.Database)
	if err != nil {
		log.Panicln(err)
	}

	i := usecase.Usecase{
		DB: db,
	}
	a := api.API{
		Cfg:        &c,
		Interactor: &i,
	}
	a.Run()

	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case s := <-term:
		// grpcServer.CatchSignal(s)
		log.Println("Exiting gracefully...", s)
	}
	return 0
}
