/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "go-cobra",
	Short: "A sample CLI for go-cobra",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:`,
}

func Execute() {

	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringP("directory", "d", "/tmp", "path")
	rootCmd.PersistentFlags().Uint("depth", 2, "Depth of search")

	viper.BindPFlag("directory", rootCmd.PersistentFlags().Lookup("directory"))
	viper.BindPFlag("depth", rootCmd.PersistentFlags().Lookup("depth"))
}
