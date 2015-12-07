package editor

import (
	"bytes"
	"errors"
)

// ErrNoEditor occurs when attempting to make an edit and no user-defined editor
// is configured.
var ErrNoEditor = errors.New("No EDITOR configured")

// ErrEditAborted is returned when the user cancels an edit (by not saving the
// temporary file in their editor)
var ErrEditAborted = errors.New("edit aborted")

// ErrUnchanged is returned when the edit did not change the input data
var ErrUnchanged = errors.New("unchanged")

// Driver represents a connection to the user's preferred text editor.
type Driver struct {
}

// ValidateFn represents a function to validate the editted content.  An edit
// will be considered a failure if an error is returned.
type ValidateFn func(input, output []byte) error

// Webhook represents the human-editable form of a webhook
type Webhook struct {
	URL               string `yaml:"url"`
	DestinationFilter string `yaml:"destination_filter"`
	MemoTypeFilter    string `yaml:"memo_type_filter"`
	MemoFilter        string `yaml:"memo_filter"`
}

// EnsureChange is a helper validation function that can be provided to `Edit`
// that ensures that the input data was changed in some way.
func EnsureChange(input, output []byte) error {
	if bytes.Equal(input, output) {
		return ErrUnchanged
	}
	return nil
}
