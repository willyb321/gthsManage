package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var idRSA string
var ip string
var port string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "gthsManage",
	Short: "Manage GTHS Noticeboard",
	Long:  `Stuff to keep the GTHS Noticeboard running.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	home, _ := homedir.Dir()
	defaultSSHPath := filepath.Join(home, ".ssh", "id_rsa")
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.gthsManage.yaml)")
	RootCmd.PersistentFlags().StringVarP(&ip, "ip", "i", "10.178.x.x", "IP of noticeboard")
	RootCmd.PersistentFlags().StringVarP(&port, "port", "p", "1471", "SSH port of noticeboard")
	RootCmd.PersistentFlags().StringVarP(&idRSA, "idfile", "f", defaultSSHPath, "Full path to your private ssh key")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Search config in home directory with name ".gthsManage" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gthsManage")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
