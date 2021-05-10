package cobra_test

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	. "go.octolab.org/toolkit/cli/cobra"
)

func TestCompletionCommand(t *testing.T) {
	tests := map[string]struct {
		format   string
		expected string
	}{
		"Bash":       {"bash", "# bash completion V2 for cli"},
		"fish":       {"fish", "# fish completion for cli"},
		"Zsh":        {"zsh", "#compdef _cli cli"},
		"PowerShell": {"powershell", "# powershell completion for cli"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			for _, cmd := range []*cobra.Command{
				NewCompletionCommand(),
				{Use: "hack for HasSubCommands"},
			} {
				buf := bytes.NewBuffer(nil)
				app := &cobra.Command{Use: "cli"}
				app.AddCommand(cmd)
				app.SetArgs([]string{"completion", test.format})
				app.SetOut(buf)

				assert.NoError(t, app.Execute())
				assert.Contains(t, buf.String(), test.expected)
			}
		})
	}
}
