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
	"crypto/rsa"
	"os"
	"encoding/pem"
	"crypto/x509"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"crypto/rand"

	"path/filepath"
	"github.com/mitchellh/go-homedir"
)

// gensshkeyCmd represents the gensshkey command
var gensshkeyCmd = &cobra.Command{
	Use:   "gensshkey",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		makeSSHKey()
		fmt.Println("gensshkey called")
	},
}

func init() {
	RootCmd.AddCommand(gensshkeyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gensshkeyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gensshkeyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func makeSSHKey() {
	dir, _ := homedir.Dir()
	pubKey := filepath.Join(dir, ".ssh", "id_rsa.pub")
	privKey := filepath.Join(dir, ".ssh", "id_rsa")
	if _, err := os.Stat(filepath.Join(dir, ".ssh")); os.IsNotExist(err) {
		os.Mkdir(filepath.Join(dir, ".ssh"), 0644)
	}
	var allowedToMake bool
	if _, err := os.Stat(privKey); os.IsNotExist(err) {
		allowedToMake = true
	} else {
		allowedToMake = false
	}
	if _, err := os.Stat(privKey); os.IsNotExist(err) {
		allowedToMake = true
	} else {
		allowedToMake = false
	}
	if allowedToMake == true {
		err := MakeSSHKeyPair(pubKey, privKey)
		if err != nil {
			fmt.Println(err)
		}
		return
	} else {
		fmt.Println("Can't make SSH Key, already exists.")
		return
	}
}

// MakeSSHKeyPair make a pair of public and private keys for SSH access.
// Public key is encoded in the format for inclusion in an OpenSSH authorized_keys file.
// Private Key generated is PEM encoded
func MakeSSHKeyPair(pubKeyPath, privateKeyPath string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return err
	}

	// generate and write private key as PEM
	privateKeyFile, err := os.Create(privateKeyPath)
	defer privateKeyFile.Close()
	if err != nil {
		return err
	}
	privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
		return err
	}

	// generate and write public key
	pub, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(pubKeyPath, ssh.MarshalAuthorizedKey(pub), 0655)
}
