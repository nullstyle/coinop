package cmd

import (
	"log"
	"net/http"
	"sync"

	uhttp "github.com/nullstyle/coinop/ui/http"
	"github.com/nullstyle/coinop/usecase"
	"github.com/spf13/cobra"
	"goji.io"
	"goji.io/pat"
	"golang.org/x/net/context"
)

var serve = &cobra.Command{
	Use:              "serve",
	Short:            "runs coinop",
	PersistentPreRun: initInjector,
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		wg.Add(1)
		go runAPI(&wg)

		wg.Add(1)
		go listenForPayments(&wg)

		wg.Wait()
	},
}

func runAPI(wg *sync.WaitGroup) {
	mux := goji.NewMux()

	// List all webhooks
	mux.HandleFuncC(
		pat.Get("/webhooks"),
		func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
			ui := uhttp.UI{ctx, r, w}
			kase := usecase.ListWebhooks{
				DB:  &drivers.DB,
				Out: &ui,
			}

			if err := kase.Exec(); err != nil {
				fail(err)
			}
		},
	)

	// Create a new webhook
	mux.HandleFuncC(
		pat.Post("/webhooks"),
		func(ctx context.Context, w http.ResponseWriter, r *http.Request) {},
	)

	// Delete a webhook
	mux.HandleFuncC(
		pat.Delete("/webhooks/:id"),
		func(ctx context.Context, w http.ResponseWriter, r *http.Request) {},
	)

	http.ListenAndServe("localhost:8000", mux)
	wg.Done()
}

func listenForPayments(wg *sync.WaitGroup) {
	kase := &usecase.WatchPayments{
		Payments: &drivers.Horizon,
		Processor: &usecase.ProcessPayment{
			Hooks:      &drivers.DB,
			Deliveries: &drivers.DB,
			Sender:     &drivers.HTTP,
		},
	}

	log.Printf("info: listening to %s", drivers.Horizon.BaseURL)

	for {
		cursor, err := drivers.DB.LoadCursor()
		if err != nil {
			log.Fatalf("error: %s", err)
		}

		err = kase.Exec(cursor)
		if err != nil {
			log.Fatalf("error: %s", err)
		}
	}
	wg.Done()
}
