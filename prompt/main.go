package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	line, err := readPassword("Password: ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("inputted password: %q\n", line)
}

func readPassword(prompt string) (string, error) {
	s, err := terminal.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	defer terminal.Restore(int(os.Stdin.Fd()), s)
	t := terminal.NewTerminal(os.Stdin, "")
	return t.ReadPassword(prompt)
}
