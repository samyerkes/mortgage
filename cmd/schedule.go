/*
Copyright © 2024 Sam Yerkes
*/
package cmd

import (
	"fmt"
	"mortgage/loan"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// scheduleCmd represents the schedule command
var scheduleCmd = &cobra.Command{
	Use:     "schedule",
	Short:   "Print an amortization schedule",
	Long:    `Print an amortization schedule for the loan, including the payment number, the amount paid, the principal paid, the interest paid, and the remaining balance.`,
	Aliases: []string{"sched", "amort", "amortization"},
	Run: func(cmd *cobra.Command, args []string) {

		var l loan.Loan

		err := viper.Unmarshal(&l)
		if err != nil {
			fmt.Println(err)
		}

		l.PrintAmortizationSchedule()
	},
}

func init() {
	rootCmd.AddCommand(scheduleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scheduleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scheduleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
