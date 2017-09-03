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
	"github.com/spf13/viper"
)

// everythingisbrokenhelpCmd represents the everythingisbrokenhelp command
var everythingisbrokenhelpCmd = &cobra.Command{
	Use:   "everythingisbrokenhelp",
	Short: "Contact me directly.",
	Long:  `Contact me directly. Only use if literally everything is broken yeah?`,
	Run: func(cmd *cobra.Command, args []string) {
		help := everythingisbrokenhelp()
		if help {
			fmt.Println("Expect a call in a bit.")
		} else {
			fmt.Println("dunno")
		}
	},
}
var phone string

func everythingisbrokenhelp() bool {
	fmt.Println("This will send a notification to me - so don't abuse it yeah. Only use if literally everything is broken")
	fmt.Println("Are you sure you want to do this? y/n")
	confirm := askForConfirmation()
	if confirm != false {
		client := sshConnect()
		session, err := client.NewSession()
		if err != nil {
			fmt.Println(err)
		}

		if viper.IsSet("phoneverify") && phone == "" {
			phone = viper.GetString("phoneverify")
		}
		err = session.Start("/home/gths/everythingisbrokenhelp.sh " + phone)
		if err != nil {
			fmt.Println(err)
		}
		return true
	}
	return false
}
func init() {
	RootCmd.AddCommand(everythingisbrokenhelpCmd)

	// Here you will define your flags and configuration settings.
	everythingisbrokenhelpCmd.Flags().StringVar(&phone, "phone", "", "The correct phone number.")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// everythingisbrokenhelpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// everythingisbrokenhelpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
