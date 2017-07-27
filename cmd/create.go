// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
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
	configToGet := []string{"IP Address", "Port", "SSH Key File"}
	for text != "q" && i < 3 { // break the loop if text == "q"
		fmt.Print("Enter the " + configToGet[i] + ": ")
		scanner.Scan()
		text = scanner.Text()
		if text != "q" {
			configs = append(configs, text)
			dir, _ := homedir.Dir()
			if i == 2 {
				data := fmt.Sprintf("ip: %v\nport: %v\nidfile: %s", configs[0], configs[1], filepath.Join(configs[2]))
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
