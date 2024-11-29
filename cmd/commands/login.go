/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package commands

import (
	"fmt"
	"heyweek/cli/pkg/models"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func NewCmdAuth() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "login",
		Short:   "Login",
		Example: "hw login",
		Run: func(cmd *cobra.Command, args []string) {
			state, err := models.NewState()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			p := tea.NewProgram(state)
			_, err = p.Run()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	return cmd
}
