package cmd

import (
	"errors"
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
	logger, err := utils.NewLogger(conf)
	if err != nil {
		return err
	}
	defer func() error {
		if logger != nil && logger.File != nil {
			_ = logger.File.Close()
		}
		if logger == nil {
			return errors.New("Logger cann't initiated")
		}
		return nil
	}()
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