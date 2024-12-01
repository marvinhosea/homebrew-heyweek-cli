/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package commands

import (
	"fmt"
	"heyweek/cli/pkg/models"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func NewCmdAuth() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "login",
		Short:   "Login",
		Example: "hw login",
		RunE: func(cmd *cobra.Command, args []string) error {
			state, err := models.NewState()
			if err != nil {
				fmt.Println(err)
				return err
			}

			p := tea.NewProgram(state)
			_, err = p.Run()
			if err != nil {
				fmt.Println(err)
				return err
			}

			return nil
		},
	}

	return cmd
}
