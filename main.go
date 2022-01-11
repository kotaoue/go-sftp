package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {

	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main() error {
	host := "localhost"
	port := "2222"
	user := "foo"
	pass := "pass"

	fmt.Println("Create sshClientConfig")
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	fmt.Println("SSH connect")
	addr := fmt.Sprintf("%s:%s", host, port)
	fmt.Println(addr)

	conn, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return err
	}

	fmt.Println("open an SFTP session over an existing ssh connection")
	client, err := sftp.NewClient(conn)
	if err != nil {
		return err
	}
	defer client.Close()

	fmt.Println("walk a directory")
	w := client.Walk("./")
	for w.Step() {
		if w.Err() != nil {
			continue
		}
		log.Println(w.Path())
	}

	fmt.Println("leave your mark")
	f, err := client.Create("./upload/hello.txt")
	if err != nil {
		return err
	}
	if _, err := f.Write([]byte("Hello world!")); err != nil {
		return err
	}
	f.Close()

	fmt.Println("check it's there")
	fi, err := client.Lstat("./upload/hello.txt")
	if err != nil {
		return err
	}
	log.Println(fi)

	return nil
}
