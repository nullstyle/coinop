package cmd

import (
	"fmt"
	"log"
	"sync"

	"github.com/nullstyle/coinop/usecase"
	"github.com/spf13/cobra"
)

var serve = &cobra.Command{
	Use:              "serve",
	Short:            "runs coinop",
	PersistentPreRun: initInjector,
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		fmt.Println("TODO: start the rest server")

		wg.Add(1)
		go listenForPayments(&wg)

		wg.Wait()
	},
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
