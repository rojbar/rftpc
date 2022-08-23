/*
Copyright © 2022 rojbar

*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"

	data "github.com/rojbar/sftpc/structs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a file to the specified channel in a server",
	Long: `Send recieves the filepath, server name and channel number to send. An example:

		send [filepath] [servername] [channelname]
	`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		filePath := args[0]
		serverName := args[1]
		//channelName := args[2]
		serverKey := "knownhosts." + serverName

		if !viper.IsSet(serverKey) {
			cobra.CheckErr(errors.New("Host name not found, add a host with add command"))
		}

		var server data.Server
		errU := viper.UnmarshalKey(serverKey, &server)
		cobra.CheckErr(errU)

		file, errO := os.Open(filePath)
		cobra.CheckErr(errO)

		defer file.Close()

		reader := bufio.NewReader(file)
		conn, err := net.Dial("tcp", server.Domain+":"+server.Port)
		cobra.CheckErr(err)
		defer conn.Close()

		writer := bufio.NewWriter(conn)

		message := make([]byte, 4096)
		bytesRead, errP := reader.Read(message)
		cobra.CheckErr(errP)
		bytesWritten, errW := writer.Write(message)
		writer.Flush()
		cobra.CheckErr(errW)

		fmt.Println(bytesRead)
		fmt.Println(bytesWritten)
		fmt.Println(string(message))

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
