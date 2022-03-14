package git

import (
	"fmt"
	"os/exec"
	"strings"
)

func RunCommand(dir string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	data, err := cmd.CombinedOutput()
	output := string(data)
	text := strings.TrimSpace(output)

	if err != nil {
		return "", fmt.Errorf(text)
	}
	return text, nil
}
