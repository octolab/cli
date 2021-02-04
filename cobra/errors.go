package cobra

import (
	"github.com/spf13/cobra"

	"go.octolab.org/toolkit/cli/errors"
)

// SilentError makes the cause error silent.
// It useful for graceful.Shutdown function.
//
//  import (
//  	"os"
//
//  	"github.com/spf13/cobra"
//  	"go.octolab.org/safe"
//  	"go.octolab.org/toolkit/cli/graceful"
//  )
//
//  cmd := cobra.Command{
//  	RunE: func(cmd *cobra.Command, args []string) error {
//  		return SilentError(cmd, action(), 2, "failed action")
//  	}
//  }
//
//  safe.Do(
//  	func() error { return cmd.ExecuteContext(ctx) },
//  	graceful.Shutdown(os.Stderr, os.Exit),
//  )
//
func SilentError(
	cmd *cobra.Command,
	cause error, code int, message ...string,
) errors.Silent {
	err := errors.NewSilent(cause, code, message...)
	if err != nil {
		cmd.SilenceErrors = true
	}
	return err
}
