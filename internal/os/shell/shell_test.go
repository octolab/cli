package shell_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "go.octolab.org/toolkit/cli/internal/os/shell"
)

func TestClassify(t *testing.T) {
	tests := map[string]struct {
		bin        string
		operations []Operation
		expected   Shell
	}{
		"sh":   {"/bin/sh", []Operation{All}, Sh},
		"bash": {"/bin/bash", []Operation{All}, Bash},
		"zsh":  {"/usr/local/bin/zsh", []Operation{All}, Zsh},
		"powershell": {
			`C:/Windows/System32/WindowsPowershell/v1.0/powershell.exe`,
			[]Operation{All},
			PowerShell,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			sh, err := Classify(test.bin, test.operations...)
			require.NoError(t, err)
			assert.Equal(t, test.expected, sh)
			assert.Equal(t, name, sh.String())
		})
	}

	t.Run("unclassified", func(t *testing.T) {
		sh, err := Classify("/usr/local/bin/fish", All)
		assert.Error(t, err)
		assert.Empty(t, sh)
		assert.Empty(t, sh.String())
	})
	t.Run("panic", func(t *testing.T) {
		assert.Panics(t, func() { _, _ = Classify("", All) })
	})
}
