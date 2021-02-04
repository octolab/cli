package errors_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "go.octolab.org/toolkit/cli/errors"
)

func TestSilent(t *testing.T) {
	type causer interface {
		Cause() error
	}

	type wrapped interface {
		Unwrap() error
	}

	t.Run("common error", func(t *testing.T) {
		cause := errors.New("a common error")
		err := NewSilent(cause, 2, "graceful exit message")

		msg, has := err.Message()

		assert.Equal(t, 2, err.Code())
		assert.True(t, has)
		assert.Equal(t, "graceful exit message", msg)
		assert.EqualError(t, err, "a common error")

		t.Run("caused", func(t *testing.T) {
			var src causer
			require.True(t, errors.As(err, &src))
			assert.True(t, errors.Is(src.Cause(), cause))
		})

		t.Run("wrapped", func(t *testing.T) {
			var src wrapped
			require.True(t, errors.As(err, &src))
			assert.True(t, errors.Is(src.Unwrap(), cause))
		})
	})

	t.Run("nil error", func(t *testing.T) {
		err := NewSilent(nil, 2, "graceful exit message")

		assert.Nil(t, err)
	})
}
