/*
Copyright © 2022 rojbar

*/
package cmd

import (
	"errors"

	data "github.com/rojbar/sftpc/structs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [domain] [port] [alias]",
	Short: "Adds a server to the known list",
	Long: `Adds a server to the know list of servers
	`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {

		key := "knownhosts." + args[2]
		if viper.IsSet(key) {
			cobra.CheckErr(errors.New("Host Name already added"))
		}

		known_hosts := make(map[string]data.Server)
		errU := viper.UnmarshalKey("knownhosts", &known_hosts)
		cobra.CheckErr(errU)

		known_hosts[args[2]] = data.Server{Domain: args[0], Port: args[1], Name: args[2]}
		viper.Set("knownhosts", known_hosts)
		viper.WriteConfig()
	},
	DisableFlagsInUseLine: true,
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
