package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var db = &cobra.Command{
	Use:              "db [command]",
	Short:            "commands to manage coinop's postgres db",
	PersistentPreRun: initInjector,
}

var dbInit = &cobra.Command{
	Use:   "init",
	Short: "install schema",
	Long:  "init initializes the postgres database used by coinop.",
	Run: func(cmd *cobra.Command, args []string) {
		log.Fatal(drivers.DB.MigrateSchema())
	},
}
