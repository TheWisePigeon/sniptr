package cmd

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Create a new snippe. Usage `sniptr set <snippet_name> <snippet>`",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("sqlite3", Sniptr_db)
		if err != nil {
			log.Fatal(err)
		}
		_ = db
		if len(args) != 2 {
			log.Fatal(`
      Missing arguments. You must provide the snippet name and content. 'sniptr set <snippet_name> <snippet_content> '
      `)
		}
	},
}
