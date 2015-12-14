package horizon

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/nullstyle/coinop/entity"
	"github.com/nullstyle/coinop/usecase"
)

// StreamPayments streams payments into `fn`, starting at `cursor` by listening
// to a horizon instance.
func (driver *Driver) StreamPayments(
	cursor string,
	fn usecase.PaymentHandler,
) error {
	log.Printf("checking for payments after %s", cursor)

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
	scanner := bufio.NewScanner(resp.Body)
	scanner.Split(splitSSE)

	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			continue
		}

		ev, err := parseEvent(scanner.Bytes())
		if err != nil {
			return err
		}

		if ev.Event != "message" {
			continue
		}

		var ent entity.Payment
		data := ev.Data.(string)
		ent, err = driver.readPayment([]byte(data))
		if err != nil {
			return err
		}

		err = fn(ent)
		if err != nil {
			return err
		}
	}

	err = scanner.Err()
	if err == io.ErrUnexpectedEOF {
		return nil
	}
	if err != nil {
		return err
	}

	return nil
}

func (driver *Driver) readPayment(data []byte) (ent entity.Payment, err error) {
	var res payment
	err = res.Unmarshal(data)
	if err != nil {
		return
	}

	// TODO: load memo data
	res.Memo.Type = "none"
	res.Memo.Value = ""

	ent, err = res.Entity()
	return
}

var _ usecase.PaymentProvider = &Driver{}
