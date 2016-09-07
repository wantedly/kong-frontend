package main

import (
	"fmt"
	"os"

	"github.com/wantedly/kong-frontend/config"
	"github.com/wantedly/kong-frontend/server"
)

func main() {
	printVersion()

	config, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}

	server.Run(config)
}
