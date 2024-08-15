/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"mortgage/loan"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:     "summary",
	Short:   "Print a summary of the loan",
	Long:    `Print a summary of the loan, including the total amount of the loan, the total interest paid, and the total amount paid.`,
	Aliases: []string{"sum", "info"},
	Run: func(cmd *cobra.Command, args []string) {
		var summary loan.Summary

		err := viper.Unmarshal(&summary)
		if err != nil {
			fmt.Println(err)
		}

		summary.Write()
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
