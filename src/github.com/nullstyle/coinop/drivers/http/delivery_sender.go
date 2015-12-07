package http

import (
	"github.com/nullstyle/coinop/entity"
	"github.com/nullstyle/coinop/usecase"
	"net/url"
)

// SendDelivery performs the deliver using the http client for this driver
func (driver *Driver) SendDelivery(d entity.Delivery) error {
	resp, err := driver.HTTP.PostForm(d.URL.String(), url.Values{
		"id":   {d.Payment.PagingToken},
		"from": {string(d.Payment.From)},
		"to":   {string(d.Payment.To)},
		//TODO:  add the other fields
	})
	if err != nil {
		return err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return nil
	}

	return &DeliveryError{
		URL:  d.URL.String(),
		Resp: resp,
	}
}

var _ usecase.DeliverySender = &Driver{}
