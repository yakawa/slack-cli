package cmd

import (
	"errors"
	"fmt"
	"os"
	"runtime"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
	"github.com/yakawa/slack-cli/cmd/file"
	"github.com/yakawa/slack-cli/cmd/msg"
)

var (
	cui     = rwi.New()
	cfgFile = ""
)

func newRootCmd(ui *rwi.RWI, args []string) *cobra.Command {
	cui = ui
	rootCmd := &cobra.Command{
		Use:   "slack-cli",
		Short: "Slack tools for Command Line",
		Long:  "Slack tools for Command Line",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("no command")
		},
	}

	rootCmd.SetArgs(args)
	rootCmd.SetOutput(ui.ErrorWriter())

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config File")

	rootCmd.AddCommand(newCheckAuthCmd())

	rootCmd.AddCommand(msg.NewMsgRootCmd(ui, args))
	rootCmd.AddCommand(file.NewFileRootCmd(ui, args))

	cobra.OnInitialize(readConfig)
	return rootCmd
}

//
func Execute(ui *rwi.RWI, args []string) (exit exitcode.ExitCode) {
	defer func() {
		if r := recover(); r != nil {
			cui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, src, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				cui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ":", src, ":", line)
			}
			exit = exitcode.Abnormal
		}
	}()

	exit = exitcode.Normal

	if err := newRootCmd(ui, args).Execute(); err != nil {
		exit = exitcode.Abnormal
	}
	return exit
}

func init() {
}

func isExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func readConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		c, _ := os.Getwd()
		viper.AddConfigPath(c)
		if h, e := homedir.Dir(); e == nil {
			viper.AddConfigPath(h)
		} else {
			fmt.Println("Cannot get Home")
		}
		viper.SetConfigName(".slack-cli")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(255)
	}
}
