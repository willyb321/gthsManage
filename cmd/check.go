// Copyright © 2017 Willyb321 <wbwilliam7@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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
