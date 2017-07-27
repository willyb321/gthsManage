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
