package cobra_test

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	. "go.octolab.org/toolkit/cli/cobra"
)

func TestCompletionCommand(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		expected string
	}{
		{"Bash", "bash", "# bash completion for cli"},
		{"PowerShell", "powershell", "Register-ArgumentCompleter -Native -CommandName 'cli' -ScriptBlock"},
		{"Zsh", "zsh", "#compdef _cli cli"},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			app := &cobra.Command{Use: "cli"}
			app.AddCommand(NewCompletionCommand())
			app.SetArgs([]string{"completion", tc.format})
			app.SetOut(buf)

			assert.NoError(t, app.Execute())
			assert.Contains(t, buf.String(), tc.expected)
		})
	}
}
