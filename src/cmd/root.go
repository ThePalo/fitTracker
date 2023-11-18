package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fitTracker",
	Short: "fitTracker App",
	Long:  `fitTracker App gives all the infor from fitTracker`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Errorf("unexpected error executing root cmd %v", err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// add flags here
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
}
