package executor

import (
	"bytes"
	"context"
	"fmt"
	osexec "os/exec"
	"strings"
)

// CommandExecutor implements the Runner interface using os/exec.
type CommandExecutor struct{}

// New creates a new CommandExecutor.
func New() Runner {
	return &CommandExecutor{}
}

// Run executes a command and returns its trimmed stdout.
func (c *CommandExecutor) Run(ctx context.Context, name string, args ...string) (string, error) {
	cmd := osexec.CommandContext(ctx, name, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%w: %s", err, strings.TrimSpace(stderr.String()))
	}

	return strings.TrimSpace(stdout.String()), nil
}
