package usecase

import (
	"fmt"

	"github.com/nullstyle/coinop/entity"
)

// CreateWebhook creates a webhook, which will later trigger deliveries from
// matching payments that occur on a digital payment network.
type CreateWebhook struct {
	DB WebhookRepository
}

// Exec runs the use case.
func (kase *CreateWebhook) Exec(hook entity.Webhook) (uid RepoID, err error) {
	hook.ID = &RepoID{}

	err = hook.Valid()
	if err != nil {
		err = &CreateWebhookError{Step: "validate", Child: err}
		return
	}

	err = kase.DB.SaveWebhook(&hook)
	if err != nil {
		err = &CreateWebhookError{Step: "repo", Child: err}
		return
	}

	if hook.IsNew() {
		// TODO: add some further explanation that the repo failed to assign
		// and ID.
		err = &CreateWebhookError{Step: "repo"}
		return
	}

	uid = *hook.ID.(*RepoID)
	return
}

// CreateWebhookError represents a failure to save a webhook into persistent
// storage.
type CreateWebhookError struct {
	Step  string
	Child error
}

func (err *CreateWebhookError) Error() string {
	base := fmt.Sprintf("failed to create webhook (%s)", err.Step)

	if err.Child != nil {
		return fmt.Sprintf("%s: %s", base, err.Child)
	}

	return base
}
