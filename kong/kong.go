package main

import (
	"fmt"
	"os"

	"github.com/fabiorphp/kongo"
	"github.com/koudaiii/kong-oauth-token-generator/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		os.Exit(1)
	}

	kong, err := kongo.New(config.KongAdminURL)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(&kong)
}
