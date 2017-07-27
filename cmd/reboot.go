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
	"fmt"

	"github.com/spf13/cobra"
)

// rebootCmd represents the reboot command
var rebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "Reboot the GTHS Noticeboard.",
	Long:  "Reboot the GTHS Noticeboard. It will take a while to boot after running this.\nThe board will redeploy automatically when it has booted.",
	Run: func(cmd *cobra.Command, args []string) {
		rebooted := reboot()
		if rebooted {
			fmt.Println("Reboot called")
		} else {
			fmt.Println("Not rebooting.")
		}
	},
}

func reboot() bool {
	fmt.Println("It is recommended that you can see the GTHS Noticeboard before rebooting. This ensures that you can\nget the board running again if something breaks")
	fmt.Println("Are you sure you want to reboot? This will take a while. y/n")
	confirm := askForConfirmation()
	if confirm != false {
		client := sshConnect()
		session, err := client.NewSession()
		if err != nil {
			fmt.Println(err)
		}
		err = session.Start("sudo shutdown -r now")
		if err != nil {
			fmt.Println(err)
		}
		return true
	}
	return false
}

func init() {
	RootCmd.AddCommand(rebootCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rebootCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rebootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
