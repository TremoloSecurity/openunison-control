/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "openunison-control",
	Short: "Deploys OpenUnison into your cluster",
	Long:  `openunison-ctl automates the deployment of OpenUnison into your cluster.  This tool will create the appropriate Secrets and deploy the correct helm charts for you`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// namespace for openunison
var namespace string

var operatorImage string
var operatorDeployCrd bool
var operatorChart string

var orchestraChart string
var orchestraLoginPortalChart string
var pathToValuesYaml string

var addClusterChart string

var clusterManagementChart string
var pathToDbPassword string
var pathToSmtpPassword string

var pathToSateliteYaml string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "openunison", "namespace to deploy openunison into")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
