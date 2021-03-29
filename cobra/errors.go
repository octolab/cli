package cobra

import (
	"github.com/spf13/cobra"

	"go.octolab.org/toolkit/cli/errors"
)

// SilentError makes the cause error silent.
// It is supported by graceful.ExitAfter or graceful.ExitAfterContext.
//
//  import (
//  	"os"
//
//  	cli "github.com/spf13/cobra"
//  	"go.octolab.org/toolkit/cli/cobra"
//  )
//
//  cmd := cli.Command{
//  	RunE: func(cmd *cli.Command, args []string) error {
//  		return cobra.SilentError(cmd, action(), 2, "failed action")
//  	}
//  }
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
