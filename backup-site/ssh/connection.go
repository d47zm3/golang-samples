package ssh

import (
	config "backup-site/config"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"os"
)

func processError(err error) {
	fmt.Printf(err.Error())
}

func publicKey(path string) ssh.AuthMethod {
	key, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		panic(err)
	}
	return ssh.PublicKeys(signer)
}

func RunCommand(cmd string, conn *ssh.Client) {
	sess, err := conn.NewSession()
	if err != nil {
		panic(err)
	}
	defer sess.Close()
	sessStdOut, err := sess.StdoutPipe()
	if err != nil {
		panic(err)
	}
	go io.Copy(os.Stdout, sessStdOut)
	sessStderr, err := sess.StderrPipe()
	if err != nil {
		panic(err)
	}
	go io.Copy(os.Stderr, sessStderr)
	err = sess.Run(cmd)
	if err != nil {
		panic(err)
	}
}

func Connect(configPath string) (*ssh.Client, error) {

	appConfig := config.ParseConfig(configPath)
	sshConfig := &ssh.ClientConfig{
		User: appConfig.Remote.Ssh.Username,
		Auth: []ssh.AuthMethod{
			publicKey(appConfig.Remote.Ssh.KeyPath),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", appConfig.Remote.Ssh.Hostname, sshConfig)
	if err != nil {
		fmt.Printf(err.Error())
	}

	return conn, err
}
