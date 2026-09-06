package main

import (
    "database/sql"
    "fmt"
    "os"
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
    "strings"
)

func main() {
    godotenv.Load()
    dsn := os.Getenv("DATABASE_URL")
    if strings.HasPrefix(dsn, "postgresql://") {
        dsn = "postgres://" + strings.TrimPrefix(dsn, "postgresql://")
    }
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
