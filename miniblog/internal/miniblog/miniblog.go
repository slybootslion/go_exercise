package miniblog

import (
	"encoding/json"
	"fmt"
	"miniblog/internal/pkg/log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func NewMinBlogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "miniblog",
		Short:        "a good go practical project",
		Long:         "a good go practical project, used to create user with basic information.",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Init(logOptions())
			defer log.Sync()
			return run()
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any argumnets, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}
	cobra.OnInitialize(initConfig)
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "the path to the miniblog configuration file. Empty string for no configuration file.")
	cmd.Flags().BoolP("toggle", "t", false, "help message for toggle")
	return cmd
}

func run() error {
	settings, _ := json.Marshal(viper.AllSettings())
	//fmt.Println(string(settings))
	//fmt.Println(viper.GetString("db.username"))
	log.Infow(string(settings))
	log.Infow(viper.GetString("db.username"))
	fmt.Println("Hello MiniBlog!!")
	return nil
}
