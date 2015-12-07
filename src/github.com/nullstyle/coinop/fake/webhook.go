package fake

import (
	"net/url"

	"github.com/icrowley/fake"
	"github.com/nullstyle/coinop/entity"
	"github.com/stellar/go-stellar-base/keypair"
)

// WebhookEntity returns a new random valid webhook
func WebhookEntity() (result entity.Webhook) {
	var err error
	result.URL, err = url.Parse("http://" + fake.DomainName())
	if err != nil {
		panic(err)
	}

	kp, err := keypair.Random()
	if err != nil {
		panic(err)
	}

	result.DestinationFilter = kp.Address()
	return
}
