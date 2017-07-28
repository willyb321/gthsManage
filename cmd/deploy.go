package cmd

import (
	"bufio"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Redeploy the noticeboard if Chrome is already open.",
	Long:  `Redeploy the noticeboard if Chrome is already open. Use deploy --fresh if not already open.`,
	Run: func(cmd *cobra.Command, args []string) {
		output, err := deploy()
		if err != nil {
			fmt.Println(err)
		}
		scanner := bufio.NewScanner(output)
		scanner.Split(bufio.ScanRunes)
		for scanner.Scan() {
			fmt.Printf(scanner.Text())
		}
		fmt.Println("deploy called")
	},
}

var fresh bool

func deploy() (output io.Reader, err error) {
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
	if fresh == false {
		err = session.Start("/home/gths/update.sh")
	} else {
		err = session.Start("/home/gths/bootboard.sh")
	}
	// client.Close()
	if err != nil {
		fmt.Println(err)
	}
	return output, err
}

func init() {
	RootCmd.AddCommand(deployCmd)
	deployCmd.Flags().BoolVar(&fresh, "fresh", false, "True if chrome not already open")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deployCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deployCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
