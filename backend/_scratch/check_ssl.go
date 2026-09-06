package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    conn, err := net.DialTimeout("tcp", "dpg-d7mgv21o3t8c73e67tu0-a.ohio-postgres.render.com:5432", 10*time.Second)
    if err != nil {
        fmt.Println("dial error:", err)
        return
    }
    defer conn.Close()
    conn.SetDeadline(time.Now().Add(10 * time.Second))
    // PostgreSQL SSLRequest message
    buf := []byte{0, 0, 0, 8, 4, 210, 22, 47}
    n, err := conn.Write(buf)
    if err != nil {
        fmt.Println("write error:", err)
        return
    }
    fmt.Println("wrote bytes:", n)
    resp := make([]byte, 1)
    n, err = conn.Read(resp)
    if err != nil {
        fmt.Println("read error:", err)
        return
    }
    fmt.Printf("response: %q (%d)\n", resp[0], n)
}
