package usecase

import (
	"github.com/nullstyle/coinop/entity"
	"golang.org/x/net/context"
	"time"
)

// ProcessDelivery attempts to perform the delivery/
//
// Steps:
// 1. Ensure delivery is valid
// 2. Lock delivery for 60 seconds
// 3. Ensure not delivered
// 4. Perform delivery, get result status
// 5. Save result status
// 6. Release Lock
type ProcessDelivery struct {
	DB     DeliveryRepository
	Sender DeliverySender
}

// Exec runs the use case.
func (kase *ProcessDelivery) Exec(
	ctx context.Context,
	d entity.Delivery,
) (err error) {

	var (
		token    int64
		deadline time.Time
	)

	token, deadline, err = kase.DB.StartDelivery(d)
	if err != nil {
		return
	}

	var cancel func()
	ctx, cancel = context.WithDeadline(ctx, deadline)
	defer cancel()

	select {
	case err = <-kase.performDelivery(d):
		if err != nil {
			// TODO: write warning if mark failed
			kase.DB.MarkDeliveryFailed(token, d)
			return err
		}

		return kase.DB.MarkDeliverySuccess(token, d)
	case <-ctx.Done():
		// our deadline expired or we were canceled, so abandon the delivery
		return
	}
}

// performDelivery asynchronously sends the delivery over the configured sender.
func (kase *ProcessDelivery) performDelivery(d entity.Delivery) <-chan error {
	result := make(chan error, 1)

	go func() {
		result <- kase.Sender.SendDelivery(d)
		close(result)
	}()

	return result
}
