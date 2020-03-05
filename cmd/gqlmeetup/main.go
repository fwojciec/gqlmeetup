package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/fwojciec/gqlmeetup/bcrypt"
	"github.com/fwojciec/gqlmeetup/dataloaden"
	"github.com/fwojciec/gqlmeetup/gqlgen"
	"github.com/fwojciec/gqlmeetup/http"
	"github.com/fwojciec/gqlmeetup/jwt"
	"github.com/fwojciec/gqlmeetup/postgres"
	"github.com/oklog/run"
)

const (
	port         = 4000
	dbConnString = "dbname=gqlmeetup sslmode=disable"
	tokenSecret  = "1p3HG04RBsJlYeG43TabTfhQC3gaIFGY"
)

func main() {
	// open database connection
	db, err := postgres.Open(dbConnString)
	if err != nil {
		logAndExit(err)
	}
	defer db.Close()

	// init resolver repository
	r := &postgres.Repository{DB: db}

	// init dataloader repository
	dlr := &postgres.DataLoaderRepository{DB: db}

	// init password service with default cost
	ps := &bcrypt.PasswordService{}

	// init dataloader service
	dls := &dataloaden.DataLoaderService{Repository: dlr}

	// init token service with default values
	ts := &jwt.TokenService{Secret: []byte(tokenSecret)}

	// init the root resolver
	resolver := &gqlgen.Resolver{
		Repository:  r,
		DataLoaders: dls,
		Password:    ps,
		Tokens:      ts,
	}

	// run things!
	var g run.Group
	{
		ln, _ := net.Listen("tcp", fmt.Sprintf(":%d", port))
		server := &http.Server{
			QueryHandler:      gqlgen.NewQueryHandler(resolver),
			PlaygroundHandler: gqlgen.NewPlaygroundHandler(),
			TokenService:      ts,
			DataLoaderService: dls,
		}
		g.Add(func() error {
			return server.Run(ln)
		}, func(error) {
			ln.Close()
		})
	}
	{
		cancel := make(chan struct{})
		g.Add(func() error {
			return interrupt(cancel)
		}, func(error) {
			close(cancel)
		})
	}
	fmt.Fprintf(os.Stdout, "🚀 Server ready at http://localhost:%d\n", port)
	if err := g.Run(); err != nil {
		logAndExit(err)
	}
}

// copied from https://github.com/oklog/oklog
func interrupt(cancel <-chan struct{}) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	select {
	case sig := <-c:
		return fmt.Errorf("received signal %s", sig)
	case <-cancel:
		return fmt.Errorf("canceled")
	}
}

func logAndExit(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
