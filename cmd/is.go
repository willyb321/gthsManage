package cmd

import (
	"github.com/spf13/cobra"
)

// isCmd represents the is command
var isCmd = &cobra.Command{
	Use:   "is",
	Short: "Checks the status of various things.",
	Long:  `Checks the status of various things. They include reboot required, deployed or not, etc`,
}

func init() {
	RootCmd.AddCommand(isCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// isCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// isCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
