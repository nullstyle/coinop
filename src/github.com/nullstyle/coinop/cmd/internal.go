package cmd

import (
	"fmt"
	"os"

	"github.com/facebookgo/inject"
	"github.com/jmoiron/sqlx"
	// postgres driver
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var injector inject.Graph

func fail(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

func init() {
	flags := Root.PersistentFlags()
	flags.String("horizon-url", "", "horizon server to use")
	flags.String("postgres-url", "", "postgres db to use")
	viper.BindPFlag("horizon-url", flags.Lookup("horizon-url"))
	viper.BindPFlag("postgres-url", flags.Lookup("postgres-url"))

	// build command tree
	Root.AddCommand(db)
	Root.AddCommand(account)
	db.AddCommand(dbInit)
	account.AddCommand(openAccount)
	account.AddCommand(closeAccount)

}

func initInjector(cmd *cobra.Command, args []string) {
	db, err := sqlx.Connect("postgres", viper.GetString("postgres-url"))
	if err != nil {
		fail(err)
	}

	mustProvide(inject.Object{Value: db})
	mustProvide(inject.Object{
		Name:  "postgres-url",
		Value: viper.GetString("postgres-url"),
	})
}

func mustProvide(val inject.Object) {
	if err := injector.Provide(&val); err != nil {
		fail(err)
	}
}
