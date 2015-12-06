package editor

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
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

// Edit activates the user's configured editor to edit the input data.  A
// temporary file with name `filename` is created an launched into the user's
// editor.
func (ed *Driver) Edit(
	input []byte,
	filename string,
	validate ValidateFn,
) (output []byte, err error) {

	// create tempdirectory
	dir, err := ioutil.TempDir("", "edit-string")
	if err != nil {
		return
	}
	defer os.RemoveAll(dir)

	// write input data to disk
	path := filepath.Join(dir, filename)
	err = ioutil.WriteFile(path, input, 0600)
	if err != nil {
		return
	}

	preStat, err := os.Stat(path)
	if err != nil {
		return
	}

	cmd, err := getEditorCommand(path)
	if err != nil {
		return
	}

	// Edit the file
	err = cmd.Run()
	if err != nil {
		return
	}

	postStat, err := os.Stat(path)
	if err != nil {
		return
	}

	if preStat.ModTime() == postStat.ModTime() {
		err = ErrEditAborted
		return
	}

	output, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = validate(input, output)
	if err != nil {
		return
	}

	return
}

// ValidateFn represents a function to validate the editted content.  An edit
// will be considered a failure if an error is returned.
type ValidateFn func(input, output []byte) error

// EnsureChange is a helper validation function that can be provided to `Edit`
// that ensures that the input data was changed in some way.
func EnsureChange(input, output []byte) error {
	if bytes.Equal(input, output) {
		return ErrUnchanged
	}
	return nil
}
