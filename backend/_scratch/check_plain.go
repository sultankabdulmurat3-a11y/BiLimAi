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
    // Basic startup message: length 8+ payload, protocol 196608, user=foo
    payload := "user=test\x00database=test\x00\x00"
    length := 4 + 4 + len(payload)
    buf := make([]byte, 4+4+len(payload))
    buf[0] = byte(length >> 24)
    buf[1] = byte(length >> 16)
    buf[2] = byte(length >> 8)
    buf[3] = byte(length)
    buf[4] = 0
    buf[5] = 3
    buf[6] = 0
    buf[7] = 0
    copy(buf[8:], []byte(payload))
    n, err := conn.Write(buf)
    if err != nil {
        fmt.Println("write error:", err)
        return
    }
    fmt.Println("wrote bytes:", n)
    resp := make([]byte, 1024)
    n, err = conn.Read(resp)
    if err != nil {
        fmt.Println("read error:", err)
        return
    }
    fmt.Printf("response bytes: %d %q\n", n, resp[:n])
}
