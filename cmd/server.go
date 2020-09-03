package main

import (
	"github.com/Marcantelj/databaseengine/pkg/databaseengine"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	srv := databaseengine.Server{
	}

	return srv.Run()
}