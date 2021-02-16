package main

import (
	ssh "backup-site/ssh"
	"fmt"
)

func main() {
	sshConnection, err := ssh.Connect("config.yaml")
	if err != nil {
		fmt.Printf("Could Not Initialise SSH Connection Due To An Error: %s", err.Error())
	}

	defer sshConnection.Close()
	ssh.RunCommand("/usr/bin/id", sshConnection)
}
