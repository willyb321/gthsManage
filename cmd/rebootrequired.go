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
