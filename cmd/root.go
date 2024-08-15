/*
Copyright © 2022 Sam Yerkes
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	Defaults = map[string]float64{
		"original_balance": 100000,
		"original_term":    12,
		"rate":             0.05,
		"escrow":           0.0,
		"additional":       0.0,
		"current_balance":  100000,
	}

	DefaultPaths = []string{
		"$HOME/.config/mortgage",
		".",
	}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "mortgage",
	Long: `Print information about a mortgage loan.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/mortgage/config.toml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	switch {
	case cfgFile != "":
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	default:
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		for k, v := range Defaults {
			viper.SetDefault(k, v)
		}

		// Search config in home directory with name "config.toml" (without extension).
		viper.AddConfigPath(home)
		for _, path := range DefaultPaths {
			viper.AddConfigPath(path)
		}
		viper.SetConfigType("toml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Unable to find config file, using defaults")
	}

}
