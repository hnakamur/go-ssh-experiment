package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	keyPair, err := keyPair("/home/hnakamur/.ssh/container_to_lxdhost.id_ed25519")
	if err != nil {
		return err
	}

	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{keyPair},
	}

	client, err := ssh.Dial("tcp", "10.155.92.21:22", config)
	if err != nil {
		return err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	err = session.Run("/bin/hostname")
	if err != nil {
		return err
	}
	fmt.Print(b.String())

	return nil
}

func keyPair(keyFile string) (ssh.AuthMethod, error) {
	pem, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}
	signer, err := ssh.ParsePrivateKey(pem)
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeys(signer), nil
}
