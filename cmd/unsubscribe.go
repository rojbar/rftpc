/*
Copyright © 2022 rojbar

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// unsubscribeCmd represents the unsubscribe command
var unsubscribeCmd = &cobra.Command{
	Use:   "unsubscribe [servername] [channelname]",
	Short: "Unsubscribes from a current server",
	Long: `Unsubscribes from a current server
		unsubscribe [servername] [channelname]
	`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("unsubscribe called")
	},
}

func init() {
	rootCmd.AddCommand(unsubscribeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unsubscribeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unsubscribeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
