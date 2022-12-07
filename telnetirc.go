package main

import (
  "bufio"
  "fmt"
  "log"
  "net"
)

func main() {
  listener, err := net.Listen("tcp", ":6667")
  if err != nil {
    log.Fatal(err)
  }
  defer listener.Close()

  log.Println("Listening on :6667")

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Println(err)
      continue
    }
    go handleConnection(conn)
  }
}

func handleConnection(conn net.Conn) {
  username := conn.RemoteAddr().String()
  log.Println("New client connected:", username)
  fmt.Fprintf(conn, "Welcome, %s!\n", username)

  scanner := bufio.NewScanner(conn)
  for scanner.Scan() {
    line := scanner.Text()
    log.Println(username, ">", line)
    fmt.Fprintf(conn, "%s said: %s\n", username, line)
  }

  if err := scanner.Err(); err != nil {
    log.Println("Connection closed:", err)
  }
  conn.Close()
}
