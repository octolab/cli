package cobra_test

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "go.octolab.org/toolkit/cli/cobra"
)

func TestCompletionCommand(t *testing.T) {
	tests := map[string]struct {
		format   string
		expected string
	}{
		"Bash":       {"bash", "# bash completion for cli"},
		"fish":       {"fish", "# fish completion for cli"},
		"Zsh":        {"zsh", "#compdef _cli cli"},
		"PowerShell": {"powershell", "Register-ArgumentCompleter -Native -CommandName 'cli' -ScriptBlock"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			app := &cobra.Command{Use: "cli"}
			app.AddCommand(NewCompletionCommand())
			app.SetArgs([]string{"completion", test.format})
			app.SetOut(buf)

			assert.NoError(t, app.Execute())
			assert.Contains(t, buf.String(), test.expected)
		})
	}

	t.Run("unreachable", func(t *testing.T) {
		command := NewCompletionCommand()
		require.Panics(t, func() { require.NoError(t, command.RunE(command, []string{"unknown"})) })
	})
}
