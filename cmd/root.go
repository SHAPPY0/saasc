package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/shappy0/saasc/internal/config"
	"github.com/shappy0/saasc/internal/core"
	"github.com/shappy0/saasc/internal/utils"
)

var (
	Version, Commit = "v0.1", "NA"
	rootCmd = &cobra.Command{
		Use:	config.AppName,
		Short:	config.ShortDesc,
		Long:	config.LongDesc,
		RunE:	Init,
	}
)

func Run() {
	if err := rootCmd.Execute(); err != nil {

	}
}

func Init(cmd *cobra.Command, args []string) error {
	conf, err := config.NewConfig().Load()
	if err != nil {
		return err
	}
	conf.Version = Version
	conf.Commit = Commit
	logFile, err := os.OpenFile(conf.LogDirPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, utils.DefaultFileMod)
	if err != nil {
		return err
	}
	defer func() {
		if logFile != nil {
			_ = logFile.Close()
		}
	}()
	var logger = utils.NewLogger(conf.LogLevel, logFile)
	app, err := core.NewApp(conf, logger)
	if err != nil {
		return err
	}
	if err := app.Init(); err != nil {
		return err
	}
	app.RunX()
	return nil
}