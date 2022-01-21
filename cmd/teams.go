package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var teamsCmd = &cobra.Command{
	Use:   "teams",
	Short: "\nManage teams in Civo",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := cmd.Help()
		if err != nil {
			return err
		}
		return errors.New("a valid subcommand is required")
	},
}

func init() {
	rootCmd.AddCommand(teamsCmd)
	teamsCmd.AddCommand(teamsListCmd)
	teamsCmd.AddCommand(teamsCreateCmd)
	teamsCmd.AddCommand(teamsRenameCmd)
	teamsCmd.AddCommand(teamsDeleteCmd)

	teamsRenameCmd.Flags().StringVarP(&newTeamName, "new-name", "n", "", "the new name for the team.")
}
