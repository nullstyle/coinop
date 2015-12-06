package editor

import (
	"os"
	"os/exec"
	"strings"
)

func getEditorCommand(path string) (*exec.Cmd, error) {
	env := os.Getenv("EDITOR")
	if env == "" {
		return nil, ErrNoEditor
	}

	pieces := strings.Split(env, " ")
	pieces = append(pieces, path)

	return exec.Command(pieces[0], pieces[1:]...), nil
}
