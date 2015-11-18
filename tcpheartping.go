package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"
)

func main() {
	args := os.Args[1:]
	argslen := len(args)
	if argslen < 2 {
		fmt.Println("tcpheartping   127.0.0.1 8080")
		os.Exit(0)
	}
	host := args[0]
	port := args[1]
	longpoll := 1
	for longpoll == 1 {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", host, port), time.Second)
		if err != nil {
			fmt.Println("Service not available, Try to restart ssh service")
			restartService()
		} else {
			_ = conn.SetDeadline(time.Now().Add(time.Second * 15)) //Set dead line for 15 seconds
			fmt.Fprintf(conn, "GET /\r\n\r\n")
			status, err := bufio.NewReader(conn).ReadString('\n')
			fmt.Println(status)
			if err != nil {
				fmt.Println(err)
				restartService()
			}
			fmt.Println("ok")
		}
		time.Sleep(3 * time.Second)
	}
}

func restartService() {
	cmd := exec.Command("systemctl", "restart", "sshpx")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(3 * time.Second) //Sleep 3 seconds for service restart
	fmt.Println("Finished: \n", out.String())

}
