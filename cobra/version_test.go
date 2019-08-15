package cobra_test

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	. "go.octolab.org/toolkit/cli/cobra"
)

func TestVersionCommand(t *testing.T) {
	tests := []struct {
		name                string
		release, date, hash string
		features            []string
	}{
		{
			"stable version",
			"1.0.0",
			"2019-07-17T12:44:00Z",
			"4f8c7f4",
			[]string{"featureA=true", "featureB=false"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
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
			for _, feature := range test.features {
				assert.Contains(t, buf.String(), feature)
			}
		})
	}
}
