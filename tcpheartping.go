package main

import (
    "fmt"
    "net"
    "bufio"
    "time"
    "os"
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
    for longpoll == 1  {
        conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", host, port))
        if err != nil {
            fmt.Println("Service not available")
            time.Sleep(10*time.Second)
            continue
        } else {
            fmt.Fprintf(conn, "GET /\r\n\r\n")
            status, err := bufio.NewReader(conn).ReadString('\n')
            fmt.Println(status)
            if err != nil {
                fmt.Println(err)
            }
            fmt.Println("ok")
        }
        time.Sleep(10*time.Second)
    }
}
