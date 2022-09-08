/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"

	data "github.com/rojbar/rftpc/structs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove [servername]",
	Short: "Removes the given server from the list of known hosts",
	Long: `Removes the given server from the list of known hosts
			remove [servername]
	`,
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {

		key := "knownhosts." + args[0]
		if !viper.IsSet(key) {
			cobra.CheckErr(errors.New("server alias not found"))
		}

		known_hosts := make(map[string]data.Server)
		errU := viper.UnmarshalKey("knownhosts", &known_hosts)
		cobra.CheckErr(errU)

		delete(known_hosts, args[0])
		viper.Set("knownhosts", known_hosts)
		errWrite := viper.WriteConfig()
		cobra.CheckErr(errWrite)

		fmt.Println("server removed successfully")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
