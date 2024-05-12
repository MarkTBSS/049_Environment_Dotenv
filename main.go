package main

import (
	"fmt"
	"os"

	"github.com/MarkTBSS/049_Environment_Dotenv/config"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env.dev"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())
	fmt.Println(cfg.App())
	fmt.Println(cfg.Db())
	fmt.Println(cfg.Jwt())
}
