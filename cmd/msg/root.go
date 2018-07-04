package msg

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

var (
	cui = rwi.New()
)

func NewMsgRootCmd(ui *rwi.RWI, args []string) *cobra.Command {
	cui = ui
	msgRootCmd := &cobra.Command{
		Use:   "msg",
		Short: "Slack tools for Command Line",
		Long:  "Slack tools for Command Line",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("no command")
		},
	}

	msgRootCmd.SetArgs(args)
	msgRootCmd.SetOutput(ui.ErrorWriter())

	msgRootCmd.AddCommand(newMsgPostCmd())

	return msgRootCmd
}
