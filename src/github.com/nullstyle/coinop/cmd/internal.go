package cmd

import (
	"fmt"
	"os"

	"github.com/facebookgo/inject"
	"github.com/jmoiron/sqlx"
	// postgres driver
	_ "github.com/lib/pq"
	"github.com/nullstyle/coinop/drivers/console"
	"github.com/nullstyle/coinop/drivers/editor"
	"github.com/nullstyle/coinop/drivers/horizon"
	"github.com/nullstyle/coinop/drivers/postgres"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var injector inject.Graph
var drivers struct {
	DB      postgres.Driver
	Horizon horizon.Driver
	Editor  editor.Driver
	Console console.Driver
}

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
	Root.AddCommand(webhook)
	db.AddCommand(dbInit)
	webhook.AddCommand(createWebhook)
	webhook.AddCommand(destroyWebhook)
	webhook.AddCommand(editWebhook)
	webhook.AddCommand(listWebhooks)

}

func initInjector(cmd *cobra.Command, args []string) {
	db, err := sqlx.Connect("postgres", viper.GetString("postgres-url"))
	if err != nil {
		fail(err)
	}

	mustProvide(inject.Object{Value: &drivers.DB})
	mustProvide(inject.Object{Value: &drivers.Horizon})
	mustProvide(inject.Object{Value: &drivers.Console})
	mustProvide(inject.Object{Value: &drivers.Editor})
	mustProvide(inject.Object{Value: db})
	mustProvide(inject.Object{
		Name:  "postgres-url",
		Value: viper.GetString("postgres-url"),
	})
	mustProvide(inject.Object{
		Name:  "horizon-url",
		Value: viper.GetString("horizon-url"),
	})

	if err := injector.Populate(); err != nil {
		fail(err)
	}
}

func mustProvide(val inject.Object) {
	if err := injector.Provide(&val); err != nil {
		fail(err)
	}
}
