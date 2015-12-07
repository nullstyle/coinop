package editor

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

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

// EditWebhook opens an editable form of `hook` in the user's editor, and
// applies the edit back over  `hook`.
func (ed *Driver) EditWebhook(hook *Webhook) error {
	buf, err := yaml.Marshal(*hook)
	if err != nil {
		return err
	}

	buf, err = ed.Edit(buf, "webhook.yaml", ed.validateWebhook)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, hook)
	if err != nil {
		return err
	}

	return nil
}

func (ed *Driver) validateWebhook(input, output []byte) error {
	var hook Webhook
	err := yaml.Unmarshal(output, &hook)
	if err != nil {
		return err
	}

	return nil
}
