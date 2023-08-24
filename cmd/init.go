package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new snippets database",
	Run: func(cmd *cobra.Command, args []string) {
		println("Initializing new snippet database")
		home_dir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		sniptr_db := home_dir + "/.sniptr.db"
		if _, err := os.Stat(sniptr_db); err == nil {
			var confirmation string
			println("Detected old database. Do you really want to start from scratch? You will lose all your current saved snippets (y/n)")
			fmt.Scanln(&confirmation)
			if strings.ToLower(confirmation) == "y" {
				println("About to create yo shit")
			} else if strings.ToLower(confirmation) == "n" {
				println("Database initialization cancelled")
				return
			} else {
				println("Invalid choice. Aborting operation")
				return
			}
			println("Deleting current database...")
			err := os.Remove(sniptr_db)
			if err != nil {
				log.Fatal(err)
			}
			println("Creating new database...")
			_, err = os.Create(sniptr_db)
			if err != nil {
				log.Fatal(err)
			}
      println("New snippets database initialized")
      return
		} else if os.IsNotExist(err) {
			_, err := os.Create(sniptr_db)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	},
}
