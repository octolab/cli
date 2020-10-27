package cobra_test

import (
	"bytes"
	"os"
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

	t.Run("autodetect", func(t *testing.T) {
		binaries := map[string]string{
			"Bash":       "/bin/bash",
			"Zsh":        "/usr/local/bin/zsh",
			"PowerShell": "/usr/local/bin/powershell",
			"fish":       "/usr/local/bin/fish",
			"csh":        "/bin/csh",
			"sh":         "/bin/sh",
		}
		defer func(before string) { _ = os.Setenv("SHELL", before) }(os.Getenv("SHELL"))

		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				buf := bytes.NewBuffer(nil)
				app := &cobra.Command{Use: "cli"}
				app.AddCommand(NewCompletionCommand())
				app.SetArgs([]string{"completion"})
				app.SetOut(buf)

				require.Contains(t, binaries, name)
				assert.NoError(t, os.Setenv("SHELL", binaries[name]))
				assert.NoError(t, app.Execute())
				assert.Contains(t, buf.String(), test.expected)
			})
		}

		t.Run("unclassified", func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			app := &cobra.Command{Use: "cli"}
			app.AddCommand(NewCompletionCommand())
			app.SetArgs([]string{"completion"})
			app.SetOut(buf)

			assert.NoError(t, os.Setenv("SHELL", binaries["csh"]))
			assert.EqualError(t, app.Execute(), `shell: cannot classify shell by "/bin/csh"`)
		})

		t.Run("unsupported", func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			app := &cobra.Command{Use: "cli"}
			app.AddCommand(NewCompletionCommand())
			app.SetArgs([]string{"completion"})
			app.SetOut(buf)

			assert.NoError(t, os.Setenv("SHELL", binaries["sh"]))
			assert.EqualError(t, app.Execute(), "completion: sh is not supported")
		})
	})
}
