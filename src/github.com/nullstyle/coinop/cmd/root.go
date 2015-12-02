package cmd

import (
	"github.com/spf13/cobra"
)

// Root represents the base command when called without any subcommands
var Root = &cobra.Command{
	Use:   "coinop",
	Short: "pay-as-you-go helper service for the stellar network",
	Long:  `Coinop is pay-as-you-go helper service for the stellar network.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}
