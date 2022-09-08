/*
Copyright © 2022 rojbar

*/
package cmd

import (
	"errors"

	data "github.com/rojbar/rftpc/structs"
	rftpci "github.com/rojbar/rftpic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// subscribeCmd represents the subscribe command
var subscribeCmd = &cobra.Command{
	Use:                   "subscribe [server alias] [channelName]",
	Short:                 "Subscribe to a channel from a server",
	Long:                  `Subscribe to a channel from a server`,
	DisableFlagsInUseLine: true,
	Args:                  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		serverName := args[0]
		channelName := args[1]
		serverKey := "knownhosts." + serverName

		if !viper.IsSet(serverKey) {
			cobra.CheckErr(errors.New("host name not found, add a host with the add command"))
		}

		var server data.Server
		errU := viper.UnmarshalKey(serverKey, &server)
		cobra.CheckErr(errU)

		errS := rftpci.Subscribe(server.Port, server.Domain, channelName)
		cobra.CheckErr(errS)
	},
}

func init() {
	rootCmd.AddCommand(subscribeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subscribeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subscribeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
