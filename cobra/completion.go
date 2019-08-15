package cobra

import "github.com/spf13/cobra"

const (
	bashFormat       = "bash"
	zshFormat        = "zsh"
	powershellFormat = "powershell"
)

// NewCompletionCommand returns a command that helps to build autocompletion.
//
//  ```sh
//  $ source <(cli completion bash)
//  #
//  # or add into .bash_profile / .zshrc
//  # if [[ -n "$(which cli)" ]]; then
//  #   source <(cli completion bash)
//  # fi
//  #
//  # or use bash-completion / zsh-completions
//  $ cli completion bash > /path/to/bash_completion.d/cli.sh
//  $ cli completion zsh  > /path/to/zsh-completions/_cli.zsh
//  ```
//
func NewCompletionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "completion",
		Short: "Print Bash or Zsh completion",
		Long:  "Print Bash or Zsh completion.",
	}
	cmd.AddCommand(
		&cobra.Command{
			Use:   bashFormat,
			Short: "Print Bash completion",
			Long:  "Print Bash completion.",
			RunE: func(cmd *cobra.Command, args []string) error {
				return root(cmd).GenBashCompletion(cmd.OutOrStdout())
			},
		},
		&cobra.Command{
			Use:   powershellFormat,
			Short: "Print PowerShell completion",
			Long:  "Print PowerShell completion.",
			RunE: func(cmd *cobra.Command, args []string) error {
				return root(cmd).GenPowerShellCompletion(cmd.OutOrStdout())
			},
		},
		&cobra.Command{
			Use:   zshFormat,
			Short: "Print Zsh completion",
			Long:  "Print Zsh completion.",
			RunE: func(cmd *cobra.Command, args []string) error {
				return root(cmd).GenZshCompletion(cmd.OutOrStdout())
			},
		},
	)
	return cmd
}
