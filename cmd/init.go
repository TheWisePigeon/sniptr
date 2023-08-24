package cmd;

import (
  "github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
  Use:"init",
  Short: "Initialize a new snippets database",
  Run: func(cmd *cobra.Command, args []string) {
    println("Init command")
  },
}
