package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get current config",
	Long:  `Get current config.`,
	Run: func(cmd *cobra.Command, args []string) {
		getConfig()
	},
}

func getConfig() {
	file, _ := os.Open(viper.ConfigFileUsed())
	scanner := bufio.NewScanner(file)
	fmt.Println("Current Config:")
	fmt.Println("======================================")
	for scanner.Scan() {
		currLine := scanner.Text()
		fmt.Println(currLine)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Println("======================================")

}

func init() {
	configCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
