package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
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
				println("Deleting current database...")
				db, err := sql.Open("sqlite3", sniptr_db)
				if err != nil {
					log.Fatal(err)
				}
				defer db.Close()
				_, err = db.Exec("drop table if exists snippets;")
				if err != nil {
					log.Fatal(err)
				}
				println("Creating new database...")
				_, err = db.Exec(`
          create table snippets(
            label text not null unique,
            value text not null
          );
        `)
				if err != nil {
					log.Fatal(err)
				}
				println("New snippets database initialized")
				return
			} else if strings.ToLower(confirmation) == "n" {
				println("Database initialization cancelled")
				return
			} else {
				log.Fatal("Invalid choice. Aborting operation")
			}
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
