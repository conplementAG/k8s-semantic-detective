package main

import (
	"github.com/conplementAG/k8s-semantic-detective/pkg/common/logging"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "semantic-detective",
	Short: "semantic-detective - the Kubernetes developer tooling",
	Long: `
Kubernetes semantic detective is a microservice for executing continous 
cluster semantic checks and exposing them in the Prometheus format for 
monitoring / alerting integration.
	
Version 0.0.1
    `,

	Version: "0.0.1",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	logging.InitializeSimpleFormat()

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.AddCommand(createDetectCommand())

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "If set logging will be verbose")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}
