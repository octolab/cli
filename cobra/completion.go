package cobra

import "github.com/spf13/cobra"

const (
	bashFormat       = "bash"
	fishFormat       = "fish"
	powershellFormat = "powershell"
	zshFormat        = "zsh"
)

// NewCompletionCommand returns a command that helps to build autocompletion.
//
//  $ cli completion bash > /path/to/bash_completion.d/cli.sh
//  $ cli completion zsh  > /path/to/zsh-completions/_cli.zsh
//
func NewCompletionCommand() *cobra.Command {
	command := cobra.Command{
		Use:   "completion",
		Short: "print Bash, fish, Zsh or PowerShell completion",
		Long:  "Print Bash, fish, Zsh or PoserShell completion.",

		Args:      cobra.ExactValidArgs(1),
		ValidArgs: []string{bashFormat, fishFormat, powershellFormat, zshFormat},

		RunE: func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case bashFormat:
				return root(cmd).GenBashCompletion(cmd.OutOrStdout())
			case fishFormat:
				return root(cmd).GenFishCompletion(cmd.OutOrStdout(), true)
			case powershellFormat:
				return root(cmd).GenPowerShellCompletion(cmd.OutOrStdout())
			case zshFormat:
				return root(cmd).GenZshCompletion(cmd.OutOrStdout())
			}
			panic("unreachable")
		},
	}
	return &command
}
