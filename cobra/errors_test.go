package cobra_test

import (
	"errors"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	. "go.octolab.org/toolkit/cli/cobra"
)

func TestSilentError(t *testing.T) {
	t.Run("common error", func(t *testing.T) {
		cmd := new(cobra.Command)
		err := SilentError(cmd, errors.New("a common error"), 2, "graceful exit message")

		msg, has := err.Message()

		assert.Equal(t, 2, err.Code())
		assert.True(t, has)
		assert.Equal(t, "graceful exit message", msg)
		assert.EqualError(t, err, "a common error")
		assert.True(t, cmd.SilenceErrors)
	})

	t.Run("nil error", func(t *testing.T) {
		cmd := new(cobra.Command)
		err := SilentError(cmd, nil, 2, "graceful exit message")

		assert.Nil(t, err)
		assert.False(t, cmd.SilenceErrors)
	})
}
