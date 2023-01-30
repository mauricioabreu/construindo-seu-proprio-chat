package main

import (
    "io"
    "log"
    "net"
)

func main() {
    addr := "localhost:8080"
    server, err := net.Listen("tcp", addr)
    if err != nil {
        log.Fatalln(err)
    }
    defer server.Close()

    log.Printf("server running on: %s", addr)

    for {
        conn, err := server.Accept()
        if err != nil {
            log.Printf("failed accepting a new connection: %s", err.Error())
            continue
        }
        log.Println("client connected")

        go func(conn net.Conn) {
            defer conn.Close()

            io.Copy(conn, conn)
        }(conn)
    }
}

