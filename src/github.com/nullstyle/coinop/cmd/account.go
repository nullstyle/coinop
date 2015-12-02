package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var account = &cobra.Command{
	Use:   "account [command]",
	Short: "commands to manage accounts",
}

var openAccount = &cobra.Command{
	Use:   "open",
	Short: "open a new account",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO: open editor to configure account")
		fmt.Println("TODO: run OpenAccount use case")
		fmt.Println("TODO: print created account's id")
	},
}

var closeAccount = &cobra.Command{
	Use:   "close [id]",
	Short: "close an account",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Usage()
			os.Exit(1)
		}

		id := args[0]
		_ = id
		fmt.Println("TODO: run CloseAccount use case")
	},
}
