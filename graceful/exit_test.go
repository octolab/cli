package graceful_test

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"

	pkg "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	cli "go.octolab.org/toolkit/cli/errors"
	. "go.octolab.org/toolkit/cli/graceful"
)

func TestExitAfterContext(t *testing.T) {
	var (
		msg       = "graceful exit message"
		empty     = func(buf *bytes.Buffer) { assert.Empty(t, buf.Len()) }
		silent    = func(buf *bytes.Buffer) { assert.Equal(t, msg, strings.TrimSpace(buf.String())) }
		recovered = func(buf *bytes.Buffer) {
			output := buf.String()
			assert.Contains(t, output, "recovered: a common error")
			assert.Contains(t, output, "---")
			assert.Contains(t, output, "unexpected panic occurred")
			assert.Contains(t, output, "go.octolab.org/safe.Do")
		}
	)

	tests := map[string]struct {
		exit   func(int)
		action func() error
		verify func(*bytes.Buffer)
	}{
		"common error": {
			exit: func(code int) { assert.Equal(t, 1, code) },
			action: func() error {
				return errors.New("a common error")
			},
			verify: empty,
		},
		"native wrapped common error": {
			exit: func(code int) { assert.Equal(t, 1, code) },
			action: func() error {
				return fmt.Errorf("wrapped: %w", errors.New("a common error"))
			},
			verify: empty,
		},
		"pkg wrapped common error": {
			exit: func(code int) { assert.Equal(t, 1, code) },
			action: func() error {
				return pkg.Wrap(errors.New("a common error"), "wrapped")
			},
			verify: empty,
		},
		"silent error": {
			exit: func(code int) { assert.Equal(t, 2, code) },
			action: func() error {
				return cli.NewSilent(errors.New("a common error"), 2, msg)
			},
			verify: silent,
		},
		"native wrapped silent error": {
			exit: func(code int) { assert.Equal(t, 2, code) },
			action: func() error {
				return fmt.Errorf("wrapped: %w", cli.NewSilent(errors.New("a common error"), 2, msg))
			},
			verify: silent,
		},
		"pkg wrapped silent error": {
			exit: func(code int) { assert.Equal(t, 2, code) },
			action: func() error {
				return pkg.Wrap(cli.NewSilent(errors.New("a common error"), 2, msg), "wrapped")
			},
			verify: silent,
		},
		"recovered common error": {
			exit: func(code int) { assert.Equal(t, 1, code) },
			action: func() error {
				panic(errors.New("a common error"))
			},
			verify: recovered,
		},
		"recovered silent error": {
			exit: func(code int) { assert.Equal(t, 1, code) },
			action: func() error {
				panic(cli.NewSilent(errors.New("a common error"), 2, msg))
			},
			verify: recovered,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			ExitAfter(test.action, buf, test.exit)
			test.verify(buf)
		})
	}
}
