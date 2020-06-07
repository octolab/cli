package cobra_test

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	. "go.octolab.org/toolkit/cli/cobra"
)

func TestVersionCommand(t *testing.T) {
	tests := map[string]struct {
		release, date, hash string
		features            []Feature
		expected            string
	}{
		"stable version": {
			"1.0.0",
			"2019-07-17T12:44:00Z",
			"4f8c7f4",
			[]Feature{
				{"featureA", true},
				{"featureB", false},
			},
			"features    : featureA=true, featureB=false",
		},
		"stable version without features": {
			"1.0.0",
			"2019-07-17T12:44:00Z",
			"4f8c7f4",
			nil,
			"features    : -",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			buf := bytes.NewBuffer(nil)
			app := &cobra.Command{Use: "cli"}
			app.AddCommand(NewVersionCommand(test.release, test.date, test.hash, test.features...))
			app.SetArgs([]string{"version"})
			app.SetOut(buf)

			assert.NoError(t, app.Execute())
			assert.Contains(t, buf.String(), app.Use)
			assert.Contains(t, buf.String(), test.release)
			assert.Contains(t, buf.String(), test.date)
			assert.Contains(t, buf.String(), test.hash)
			assert.Contains(t, buf.String(), test.expected)
		})
	}
}
