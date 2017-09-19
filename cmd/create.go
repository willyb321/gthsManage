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
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create config file based on questions",
	Long:  `Create config file based on questions`,
	Run: func(cmd *cobra.Command, args []string) {
		createConfig()
		fmt.Println("create called")
	},
}

func createConfig() {
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	i := 0
	configs := []string{}
	configToGet := []string{"IP Address", "Port", "SSH Key File", "Phone Number (for Verification)"}
	for text != "q" && i < 4 { // break the loop if text == "q"
		fmt.Print("Enter the " + configToGet[i] + ": ")
		scanner.Scan()
		text = scanner.Text()
		if text != "q" {
			configs = append(configs, text)
			dir, _ := homedir.Dir()
			if i == 3 {
				data := fmt.Sprintf("ip: %v\nport: %v\nidfile: %s\nphoneverify: %s", configs[0], configs[1], filepath.Join(configs[2]), configs[3])
				ioutil.WriteFile(filepath.Join(dir, ".gthsManage.yaml"), []byte(data), 0644)
				fmt.Println("Wrote the following config:")
				fmt.Println(data)
			}
		}
		i++
	}
}

func init() {
	configCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
