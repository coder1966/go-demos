package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	Addr := "10.100.64.104:22"
	Command := "netstat"
	Username := "zhangsr"
	Password := "88888"
	RSA := "/home/zhangub/.ssh/id_rsa"
	Timeout := 10

	outPut, err := getData(Addr, Command, Username, Password, RSA, Timeout)
	fmt.Println(string(outPut), err)
	fmt.Println("=================================================")
	fmt.Println("=================================================")
	fmt.Println("=================================================")
	fmt.Println("=================================================")
	fmt.Println("=================================================")
	outPut, err = getData(Addr, Command, Username, Password, "", Timeout)
	fmt.Println(string(outPut), err)
}

func getData(addr, command, username, password, rsa string, timeout int) ([]byte, error) {
	var config ssh.ClientConfig

	if rsa == "" {
		// use password
		config = ssh.ClientConfig{
			User: username,
			Auth: []ssh.AuthMethod{ssh.Password(password)},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
			Timeout: time.Duration(timeout) * time.Second,
		}
	} else {
		// use rsa public key
		key, err := ioutil.ReadFile(rsa)
		if err != nil {
			err = fmt.Errorf("unable to read rsa public key: %v", err)
			return nil, err
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			err = fmt.Errorf("unable to parse rsa public key: %v", err)
			return nil, err
		}
		config = ssh.ClientConfig{
			User: username,
			Auth: []ssh.AuthMethod{ssh.PublicKeys(signer)},
			HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
				return nil
			},
			Timeout: time.Duration(timeout) * time.Second,
		}
	}

	client, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		return nil, fmt.Errorf("unable to connect: %s error %v", addr, err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return nil, fmt.Errorf("ssh new session error %v", err)
	}
	defer session.Close()

	return session.Output(command)
}
