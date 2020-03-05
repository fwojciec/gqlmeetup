package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fwojciec/gqlmeetup"
	"github.com/fwojciec/gqlmeetup/bcrypt"
	"github.com/fwojciec/gqlmeetup/postgres"
)

const dbConnString = "dbname=gqlmeetup sslmode=disable"

func main() {
	var (
		email    string
		password string
		admin    bool
	)

	fs := flag.NewFlagSet("adminuser", flag.ExitOnError)
	fs.StringVar(&email, "email", "", "user's email address (required)")
	fs.StringVar(&password, "password", "", "user's password (required)")
	fs.BoolVar(&admin, "admin", false, "give user admin privileges")
	if err := fs.Parse(os.Args[1:]); err != nil {
		if err == flag.ErrHelp {
			os.Exit(1)
		}
		logAndExit(err)
	}

	if email == "" || password == "" {
		fs.PrintDefaults()
		logAndExit(fmt.Errorf("error: required params are missing"))
	}

	// open database connection
	db, err := postgres.Open(dbConnString)
	if err != nil {
		logAndExit(err)
	}
	defer db.Close()

	// init cli repository
	r := &postgres.CLIRepository{DB: db}

	// init password service with default cost
	ps := &bcrypt.PasswordService{}

	hash, err := ps.Hash(password)
	if err != nil {
		logAndExit(err)
	}
	user := gqlmeetup.User{
		Email:    email,
		Password: hash,
		Admin:    admin,
	}
	if err := r.UserCreate(user); err != nil {
		logAndExit(err)
	}
}

func logAndExit(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
