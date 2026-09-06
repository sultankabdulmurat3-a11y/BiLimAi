package main

import (
    "database/sql"
    "fmt"
    "os"
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        fmt.Println("load error:", err)
    }
    dsn := os.Getenv("DATABASE_URL")
    fmt.Println("DSN:", dsn)
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        fmt.Println("open error:", err)
        return
    }
    defer db.Close()
    err = db.Ping()
    fmt.Println("ping error:", err)
}
