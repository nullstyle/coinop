package cmd

import (
	"fmt"
	"os"

	"github.com/nullstyle/coinop/drivers/editor"
	"github.com/nullstyle/coinop/usecase"
	"github.com/spf13/cobra"
)

var webhook = &cobra.Command{
	Use:              "webhook [command]",
	Short:            "commands to manage webhooks",
	PersistentPreRun: initInjector,
}

var createWebhook = &cobra.Command{
	Use:   "new",
	Short: "create a new webhook",
	Run: func(cmd *cobra.Command, args []string) {
		kase := usecase.CreateWebhook{DB: &drivers.DB}

		editable := editor.Webhook{
			URL:               "http://localhost:8080/edit/me",
			DestinationFilter: "enter account id here",
		}
		err := drivers.Editor.EditWebhook(&editable)
		if err != nil {
			fail(err)
		}

		hook, err := editable.Entity()
		if err != nil {
			fail(err)
		}

		_, err = kase.Exec(hook)
		if err != nil {
			fail(err)
		}
	},
}

var destroyWebhook = &cobra.Command{
	Use:   "rm [id]",
	Short: "remove a webhook",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Usage()
			os.Exit(1)
		}

		id := args[0]
		_ = id
		fmt.Println("TODO: run DestroyWebhook use case")
	},
}

var editWebhook = &cobra.Command{
	Use:   "edit [id]",
	Short: "edit a webhook",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Usage()
			os.Exit(1)
		}

		id := args[0]
		_ = id
		fmt.Println("TODO: open editor to configure webhook")
		fmt.Println("TODO: run UpdateWebhook use case")
	},
}

var listWebhooks = &cobra.Command{
	Use:   "ls",
	Short: "list webhooks",
	Run: func(cmd *cobra.Command, args []string) {
		kase := usecase.ListWebhooks{
			DB:  &drivers.DB,
			Out: &uis.Console,
		}

		if err := kase.Exec(); err != nil {
			fail(err)
		}
	},
}
