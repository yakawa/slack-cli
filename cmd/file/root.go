package file

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

var (
	cui = rwi.New()
)

func NewFileRootCmd(ui *rwi.RWI, args []string) *cobra.Command {
	cui = ui
	fileRootCmd := &cobra.Command{
		Use:   "file",
		Short: "Slack tools for Command Line",
		Long:  "Slack tools for Command Line",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("no command")
		},
	}

	fileRootCmd.SetArgs(args)
	fileRootCmd.SetOutput(ui.ErrorWriter())

	fileRootCmd.AddCommand(newFileUploadCmd())

	return fileRootCmd
}
