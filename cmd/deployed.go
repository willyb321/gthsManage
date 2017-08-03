package cmd

import (
	"bufio"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// deployedCmd represents the deployed command
var deployedCmd = &cobra.Command{
	Use:   "deployed",
	Short: "Checks if the noticeboard is deployed.",
	Long: `Checks if the noticeboard is deployed.
Will say "Deployed." if deployed, and "Not deployed." if not deployed`,
	Run: func(cmd *cobra.Command, args []string) {
		output, err := isDeployed()
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

func isDeployed() (output io.Reader, err error) {
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
	err = session.Start("/home/gths/isdeployed.sh")
	if err != nil {
		fmt.Println(err)
	}
	return output, err
}

func init() {
	isCmd.AddCommand(deployedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deployedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deployedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
