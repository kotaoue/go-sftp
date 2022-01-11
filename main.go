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

	// Create sshClientConfig
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// SSH connect.
	addr := fmt.Sprintf("%s:%s", host, port)
	fmt.Println(addr)

	conn, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return err
	}

	// open an SFTP session over an existing ssh connection.
	client, err := sftp.NewClient(conn)
	if err != nil {
		return err
	}
	defer client.Close()

	// walk a directory
	w := client.Walk("./")
	for w.Step() {
		if w.Err() != nil {
			continue
		}
		log.Println(w.Path())
	}

	// leave your mark
	f, err := client.Create("hello.txt")
	if err != nil {
		return err
	}
	if _, err := f.Write([]byte("Hello world!")); err != nil {
		return err
	}
	f.Close()

	// check it's there
	fi, err := client.Lstat("hello.txt")
	if err != nil {
		return err
	}
	log.Println(fi)

	return nil
}
