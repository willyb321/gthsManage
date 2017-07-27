package cmd

import (
	"bufio"
	"fmt"
	"os"

	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check current config for errors",
	Long:  `Check current config for errors`,
	Run: func(cmd *cobra.Command, args []string) {
		wasFine := checkConfig()
		if wasFine != true {
			fmt.Println("Config Check Unsuccessful")
		} else {
			fmt.Println("Config Check Successful")
			fmt.Println("Attempting to connect to configured server to test. If it is successful it will connect then exit.\nIf it is not working it will probably hang so just ctrl + c")
			clientTest := sshConnect()
			session, err := clientTest.NewSession()
			if err != nil {
				fmt.Println(err)
			}
			output, err := session.CombinedOutput("echo Successfully connected")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(output))
			session.Close()
			clientTest.Close()
		}
	},
}

func checkConfig() bool {
	file, _ := os.Open(viper.ConfigFileUsed())
	scanner := bufio.NewScanner(file)
	fmt.Println("Current Config:")
	fmt.Println("======================================")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		currLine := scanner.Text()
		currLineVal := strings.Split(currLine, ": ")
		var isGood bool
		if len(currLineVal) > 1 && currLineVal[1] != "" {
			isGood = true
		} else {
			isGood = false
		}
		if isGood != true {
			fmt.Println("Line not good - text: " + scanner.Text())
			return false
		}
	}
	fmt.Println("======================================")
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return true
}

func init() {
	configCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
