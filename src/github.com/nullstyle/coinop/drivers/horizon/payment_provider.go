package horizon

import (
	"fmt"
	"net/http"

	"github.com/nullstyle/coinop/usecase"
)

// StreamPayments streams payments into `fn`, starting at `cursor` by listening
// to a horizon instance.
func (driver *Driver) StreamPayments(
	cursor string,
	fn usecase.PaymentHandler,
) error {

	url := fmt.Sprintf("%s/payments?cursor=%s", driver.BaseURL, cursor)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "text/event-stream")
	resp, err := driver.HTTP.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

var _ usecase.PaymentProvider = &Driver{}
