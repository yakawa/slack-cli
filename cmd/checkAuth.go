package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yakawa/slack-cli/slack"
)

func newCheckAuthCmd() *cobra.Command {
	checkAuthCmd := &cobra.Command{
		Use:   "auth",
		Short: "Check Auth.",
		RunE: func(cmd *cobra.Command, args []string) error {
			slack.CheckAuth()
			return nil
		},
	}

	return checkAuthCmd
}
