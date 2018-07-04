package msg

import (
	"github.com/spf13/cobra"
	"github.com/yakawa/slack-cli/slack"
)

func newMsgPostCmd() *cobra.Command {
	msgPostCmd := &cobra.Command{
		Use:   "post [CHANNEL] [MESSAGE]",
		Short: "Post Message to Channel or User",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			c := args[0]
			m := args[1]

			slack.PostMessage(c, m)
			return nil
		},
	}

	return msgPostCmd
}
