package main

import (
	"fmt"
	"log"

	"github.com/sainikmandal/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	if err := cfg.SetUser("sainik"); err != nil {
		log.Fatalf("failed to update user: %v", err)
	}

	// Read back updated config to verify
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("failed to read updated config: %v", err)
	}

	fmt.Println("Current User:", cfg.CurrentUserName)
}
