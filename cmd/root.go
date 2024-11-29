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
	}

	rootCmd.SetHelpTemplate(rootCmd.HelpTemplate() + "help tmpl")
	rootCmd.AddCommand(commands.NewCmdAuth())

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(fmt.Errorf("error %w", err))
		os.Exit(1)
	}
}
