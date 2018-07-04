package file

import (
	"github.com/spf13/cobra"
	"github.com/yakawa/slack-cli/slack"
)

func newFileUploadCmd() *cobra.Command {
	uploadFileCmd := &cobra.Command{
		Use:   "upload [CHANNEL] [COMMENT] [FILE]",
		Short: "Post Message to Channel or User",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			c := args[0]
			m := args[1]
			f := args[2]

			slack.UploadFile(c, m, f)
			return nil
		},
	}

	return uploadFileCmd
}
