package cobra

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	bashFormat       = "bash"
	fishFormat       = "fish"
	powershellFormat = "powershell"
	zshFormat        = "zsh"
)

// NewCompletionCommand returns a command that helps to build autocompletion.
// It supports Bash, fish, Zsh and PowerShell shells.
//
//  $ cli completion bash > /path/to/bash_completion.d/cli.sh
//  $ cli completion zsh  > /path/to/zsh-completions/_cli.zsh
//
func NewCompletionCommand() *cobra.Command {
	command := cobra.Command{
		Use:   "completion",
		Short: "print Bash, fish, Zsh or PowerShell completion",
		Long:  "Print Bash, fish, Zsh or PoserShell completion.",

		Args:      cobra.MaximumNArgs(1),
		ValidArgs: []string{bashFormat, fishFormat, powershellFormat, zshFormat},

		RunE: func(cmd *cobra.Command, args []string) error {
			format := map[string]string{
				"bash": bashFormat,
				"fish": fishFormat,
				"zsh":  zshFormat,
			}[filepath.Base(
				os.Getenv("SHELL"),
			)]
			if len(args) > 0 {
				format = args[0]
			}

			switch format {
			case fishFormat:
				return cmd.Root().GenFishCompletion(cmd.OutOrStdout(), true)
			case powershellFormat:
				return cmd.Root().GenPowerShellCompletion(cmd.OutOrStdout())
			case zshFormat:
				return cmd.Root().GenZshCompletion(cmd.OutOrStdout())
			default:
				return cmd.Root().GenBashCompletion(cmd.OutOrStdout())
			}
		},
	}
	return &command
}
