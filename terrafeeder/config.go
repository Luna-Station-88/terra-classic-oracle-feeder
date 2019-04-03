package main

import (
	"feeder/terrafeeder/rest"
	"feeder/terrafeeder/updater"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func getHistoryPath() string {
	return viper.GetString(flagHome) + "/history.db"
}

// read saved configure for cobra
func initConfig(rootCmd *cobra.Command) {
	cobra.OnInitialize(func() {
		home := viper.GetString(flagHome)

		if home != "" {
			viper.AddConfigPath(home)
		} else {
			home, err := homedir.Dir()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			viper.AddConfigPath(home)
		}
		viper.SetConfigName(".terrafeeder")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	})

	rootCmd.PersistentFlags().String(flagHome, defaultHome, "directory for config and data")
	_ = viper.BindPFlag(flagHome, rootCmd.PersistentFlags().Lookup(flagHome))
	viper.SetDefault(flagHome, defaultHome)

}

func registCommands(cmd *cobra.Command) {

	cmd.Flags().Bool(flagNoREST, false, "run without REST Api serving")
	_ = viper.BindPFlag(flagNoREST, cmd.Flags().Lookup(flagNoREST))

	// viper.SetDefault(flagNoREST, false)

	// adding task flags
	updater.RegistCommand(cmd)
	if !viper.GetBool(flagNoREST) {
		rest.RegistCommand(cmd)
	}
}
