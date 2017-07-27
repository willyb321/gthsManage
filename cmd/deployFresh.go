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
