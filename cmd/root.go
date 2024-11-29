package cmd

import (
	"fmt"
	"heyweek/cli/cmd/commands"
	"os"

	"github.com/spf13/cobra"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "hw",
		Short: "heyweek",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	rootCmd.SetHelpTemplate(rootCmd.HelpTemplate() + "help tmpl")
	rootCmd.AddCommand(commands.NewCmdAuth())

	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error running main command %v", err)
		os.Exit(1)
	}
}
