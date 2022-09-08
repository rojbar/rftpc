/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	data "github.com/rojbar/rftpc/structs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:                   "show",
	Short:                 "Shows the list of known servers",
	Long:                  `Shows the list of known servers`,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {

		known_hosts := make(map[string]data.Server)
		errU := viper.UnmarshalKey("knownhosts", &known_hosts)
		cobra.CheckErr(errU)

		for _, element := range known_hosts {
			fmt.Println("Alias:", element.Name, "\t", "Domain:", element.Domain, "\t", "Port:", element.Port, "\t")
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
