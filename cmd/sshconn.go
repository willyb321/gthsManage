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
	"io/ioutil"
	"log"
	"time"

	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
)

func sshConnect() *ssh.Client {
	if viper.IsSet("idfile") {
		idRSA = viper.GetString("idfile")
	}
	key, err := ioutil.ReadFile(idRSA)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}
	if viper.IsSet("ip") {
		ip = viper.GetString("ip")
	}
	if viper.IsSet("port") {
		port = viper.GetString("port")
	}
	config := &ssh.ClientConfig{
		User: "gths",
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	// Connect to the remote server and perform the SSH handshake.
	fmt.Println("======================================")
	fmt.Println("SSH Config:")
	fmt.Println("SSH URL: " + ip + ":" + port)
	fmt.Println("======================================")
	fmt.Println("Is this correct? If not ctrl+c in the next 1 second")
	time.Sleep(1 * time.Second)
	client, err := ssh.Dial("tcp", ip+":"+port, config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	return client
}
