package cmd

import (
	"errors"
	"github.com/hjin-me/go-utils/v2/logex"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"preject/cmd/servercmd"
)

var Version string
var CommitId string
var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "服务启动",
	Long:  `这个服务会启动一个例子`,
}

func Execute(version, commitId string) {
	Version = version
	CommitId = commitId

	logex.Init(logrus.Fields{
		"version":   Version,
		"commit_id": CommitId,
	}, true)

	//可启用多个服务
	rootCmd.AddCommand(servercmd.Cmd)
	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stderr)

	if err := rootCmd.Execute(); err != nil {
		rootCmd.PrintErr(err.Error())
		var cerr CErr
		if errors.As(err, &cerr) {
			os.Exit(cerr.ExitCode())
			return
		}
		os.Exit(1)
	}
}

type CErr interface {
	Error() string
	ExitCode() int
}
