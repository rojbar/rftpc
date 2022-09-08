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

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send [filepath] [servername] [channelname]",
	Short: "Send a file to the specified channel in a server",
	Long: `Send recieves the filepath, server name and channel number to send. An example:

		send [filepath] [servername] [channelname]
	`,
	Args:                  cobra.ExactArgs(3),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {

		filePath := args[0]
		serverName := args[1]
		channelName := args[2]
		serverKey := "knownhosts." + serverName

		if !viper.IsSet(serverKey) {
			cobra.CheckErr(errors.New("host name not found, add a host with the add command"))
		}

		var server data.Server
		errU := viper.UnmarshalKey(serverKey, &server)
		cobra.CheckErr(errU)

		errS := rftpci.SendFile(server.Port, server.Domain, channelName, filePath)
		cobra.CheckErr(errS)
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
