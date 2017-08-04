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
