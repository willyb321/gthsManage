package cmd

import (
	"bufio"
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update packages.",
	Long: `There is a few different packages that need updating.
	This command will update all of them.`,
	Run: func(cmd *cobra.Command, args []string) {
		output, err := updateAll()
		if err != nil {
			fmt.Println(err)
		}
		scanner := bufio.NewScanner(output)
		scanner.Split(bufio.ScanRunes)
		for scanner.Scan() {
			fmt.Printf(scanner.Text())
		}
		fmt.Println("Update done")
	},
}

func updateAll() (output io.Reader, err error) {
	client := sshConnect()
	session, err := client.NewSession()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("About to update all packages, this will take a while.")
	outReader, err := session.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	errReader, err := session.StderrPipe()
	if err != nil {
		fmt.Println(err)
	}
	output = io.MultiReader(outReader, errReader)
	err = session.Start("/home/gths/updateall.sh")
	if err != nil {
		fmt.Println(err)
	}
	return output, err
}

func init() {
	RootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
