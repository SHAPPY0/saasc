package cmd

import (
	// "fmt"
	"github.com/spf13/cobra"
	"github.com/shappy0/saasc/internal/config"
	"github.com/shappy0/saasc/internal/core"
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
	app, err := core.NewApp(conf)
	if err != nil {
		return err
	}
	if err := app.Init(); err != nil {
		// logger.Error(err.Error())
		return err
	}
	app.RunX()
	return nil
}