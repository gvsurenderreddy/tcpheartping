package main

import (
    "fmt"
    "net"
    "bufio"
    "time"
)

func main() {
    longpoll := 1
    for longpoll == 1  {
        conn, err := net.Dial("tcp", "127.0.0.1:9070")
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
