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
	"io"

	"github.com/spf13/cobra"
)

// rebootrequiredCmd represents the rebootrequired command
var rebootrequiredCmd = &cobra.Command{
	Use:   "rebootrequired",
	Short: "Checks if a reboot is required.",
	Long:  `Checks if packages are waiting for a reboot.`,
	Run: func(cmd *cobra.Command, args []string) {
		output, err := rebootRequired()
		if err != nil {
			fmt.Println(err)
		}
		scanner := bufio.NewScanner(output)
		scanner.Split(bufio.ScanRunes)
		for scanner.Scan() {
			fmt.Printf(scanner.Text())
		}
	},
}

func rebootRequired() (output io.Reader, err error) {
	client := sshConnect()
	session, err := client.NewSession()
	if err != nil {
		fmt.Println(err)
	}
	outReader, err := session.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	errReader, err := session.StderrPipe()
	if err != nil {
		fmt.Println(err)
	}
	output = io.MultiReader(outReader, errReader)
	err = session.Start("/home/gths/isreboot.sh")
	if err != nil {
		fmt.Println(err)
	}
	return output, err
}

func init() {
	isCmd.AddCommand(rebootrequiredCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rebootrequiredCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rebootrequiredCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
