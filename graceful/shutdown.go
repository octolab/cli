package graceful

import (
	"errors"
	"fmt"
	"io"

	pkg "go.octolab.org/errors"
	"go.octolab.org/unsafe"

	cli "go.octolab.org/toolkit/cli/errors"
)

// Shutdown returns an error handler to make a graceful shutdown.
//
//  import (
//  	"context"
//  	"os"
//
//  	"go.octolab.org/safe"
//  )
//
//  func main() {
//  	ctx, cancel := context.WithCancel(context.Background())
//  	defer cancel()
//
//  	safe.Do(
//  		func() error { return root.ExecuteContext(ctx) },
//  		Shutdown(os.Stderr, os.Exit),
//  	)
//  }
//
func Shutdown(stderr io.Writer, exit func(int)) func(error) {
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
		exit(code)
	}
}
