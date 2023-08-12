/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github/YashPimple/litmusctl/cmd/config"
	"github/YashPimple/litmusctl/cmd/create"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "litmusctl",
	Short: "An implementation of litmusctl using promptui and cobra",
	Long:  ``,
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
	rootCmd.AddCommand(config.ConfigCmd)
	rootCmd.AddCommand(create.CreateCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}