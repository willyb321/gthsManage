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
	"io"

	"github.com/spf13/cobra"
)

// deployFreshCmd represents the deployFresh command
var deployFreshCmd = &cobra.Command{
	Use:   "deployFresh",
	Short: "Deploy the noticeboard if Chrome is not already open.",
	Long:  `Deploy the noticeboard if Chrome is not already open. Use deploy if already open.`,
	Run: func(cmd *cobra.Command, args []string) {
		output, err := deployFresh()
		if err != nil {
			fmt.Println(err)
		}
		scanner := bufio.NewScanner(output)
		scanner.Split(bufio.ScanRunes)
		for scanner.Scan() {
			fmt.Printf(scanner.Text())
		}
		fmt.Println("deployFresh called")
	},
}

func deployFresh() (output io.Reader, err error) {
	client := sshConnect()
	session, err := client.NewSession()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("About to deploy, this will take a while.")
	outReader, err := session.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	errReader, err := session.StderrPipe()
	if err != nil {
		fmt.Println(err)
	}
	output = io.MultiReader(outReader, errReader)
	err = session.Start("/home/gths/bootboard.sh")
	if err != nil {
		fmt.Println(err)
	}
	return output, err
}

func init() {
	RootCmd.AddCommand(deployFreshCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deployFreshCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deployFreshCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
