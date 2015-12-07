package usecase

import (
	"fmt"

	"github.com/nullstyle/coinop/entity"
)

type CreateWebhook struct {
	DB WebhookRepository
}

// Exec runs the use case.
func (kase *CreateWebhook) Exec(hook entity.Webhook) (uid RepoID, err error) {
	hook.ID = &RepoID{}
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
