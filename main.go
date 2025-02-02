package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/vinser/qwen-opds/config"
	"github.com/vinser/qwen-opds/db"
	"github.com/vinser/qwen-opds/server"
)

func main() {
	serviceCmd := flag.String("service", "", "Control App system service: install, start, stop, restart, uninstall, status")
	reindexCmd := flag.Bool("reindex", false, "Reindex the book stock directory")
	configCmd := flag.Bool("config", false, "Create default config file")
	helpCmd := flag.Bool("help", false, "Display help")
	versionCmd := flag.Bool("version", false, "Output version information")

	flag.Parse()

	switch {
	case *helpCmd:
		fmt.Println("Usage: qwen-opds [OPTION] [data directory]")
		os.Exit(0)
	case *versionCmd:
		fmt.Println("Version: 1.0.0")
		os.Exit(0)
	case *serviceCmd != "":
		handleServiceCommand(*serviceCmd)
	case *reindexCmd:
		reindexBooks()
	case *configCmd:
		createDefaultConfig()
	default:
		startServer()
	}
}

func handleServiceCommand(action string) {
	log.Printf("Service action: %s\n", action)
}

func reindexBooks() {
	log.Println("Reindexing books...")
}

func createDefaultConfig() {
	config.CreateDefaultConfig()
}

func startServer() {
	cfg := config.LoadConfig("config/config.yml")
	db.InitDatabase(cfg.Database.DSN)
	server.StartServer(cfg.OPDS.Port)
}
