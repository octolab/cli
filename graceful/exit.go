package graceful

import (
	"errors"
	"fmt"
	"io"

	pkg "go.octolab.org/errors"
	"go.octolab.org/safe"
	"go.octolab.org/unsafe"

	cli "go.octolab.org/toolkit/cli/errors"
)

// ExitAfter exits after the action and handles any errors that occurred by it,
// in particular any panics. It supports errors.Silent to redefine an exit code.
//
//  import (
//  	"os"
//
//  	"example.com/path/to/cobra/cmd"
//  	"go.octolab.org/toolkit/cli/graceful"
//  )
//
//  func main() {
//  	root := cmd.New(os.Stderr, os.Stdout)
//  	graceful.ExitAfter(root.Execute, os.Stderr, os.Exit)
//  }
//
// For context-based actions
//
//  import (
//  	"context"
//  	"os"
//
//  	"example.com/path/to/cobra/cmd"
//  	"go.octolab.org/fn"
//  	"go.octolab.org/toolkit/cli/graceful"
//  )
//
//  func main() {
//  	ctx, cancel := context.WithCancel(context.Background())
//  	defer cancel()
//
//  	root := cmd.New(os.Stderr, os.Stdout)
//  	graceful.ExitAfter(fn.HoldContext(ctx, root.ExecuteContext), os.Stderr, os.Exit)
//  }
//
func ExitAfter(
	action func() error,
	stderr io.Writer, callback func(int),
) {
	safe.Do(action, exit(stderr, callback))
}

func exit(stderr io.Writer, callback func(int)) func(error) {
	return func(err error) {
		code := 1
		if silent := cli.Silent(nil); errors.As(err, &silent) {
			code = silent.Code()
			if message, has := silent.Message(); has {
				unsafe.DoSilent(fmt.Fprintln(stderr, message))
			}
		} else if recovered := pkg.Recovered(nil); errors.As(err, &recovered) {
			unsafe.DoSilent(fmt.Fprintf(stderr, "recovered: %+v\n", recovered.Cause()))
			unsafe.DoSilent(fmt.Fprintln(stderr, "---"))
			unsafe.DoSilent(fmt.Fprintf(stderr, "%+v\n", err))
		}
		callback(code)
	}
}
