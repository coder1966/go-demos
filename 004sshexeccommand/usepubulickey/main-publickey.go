package main

import (
	"fmt"
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

func NewConfig(keyFile, user string) (config *ssh.ClientConfig, err error) {
	key, err := ioutil.ReadFile(keyFile)
	if err != nil {
		err = fmt.Errorf("unable to read private key: %v", err)
		return
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		err = fmt.Errorf("unable to parse private key: %v", err)
		return
	}

	config = &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return
}

func Run(config *ssh.ClientConfig, host, command string, port int) (outPut []byte, err error) {
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
	if err != nil {
		err = fmt.Errorf("unable to connect: %s error %v", host, err)
		return
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		err = fmt.Errorf("ssh new session error %v", err)
		return
	}
	defer session.Close()

	outPut, err = session.Output(command)
	if err != nil {
		err = fmt.Errorf("run command %s on host %s error %v", command, host, err)
	}
	return
}

type args struct {
	config  *ssh.ClientConfig
	host    string
	command string
	port    int
}

func main() {
	config, _ := NewConfig("/home/zhangub/.ssh/id_rsa", "zhangsr")

	args := args{
		config:  config,
		host:    "10.100.64.104",
		command: "netstat",
		port:    22,
	}

	gotOutPut, err := Run(args.config, args.host, args.command, args.port)
	if err != nil {
		fmt.Println(" error: ", err)
	}
	fmt.Println(string(gotOutPut))

}
