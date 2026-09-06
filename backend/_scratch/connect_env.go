package main

import (
    "fmt"
    "os"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    fmt.Println("load error:", err)
    fmt.Println("JWT_SECRET:", os.Getenv("JWT_SECRET"))
    fmt.Println("DATABASE_URL:", os.Getenv("DATABASE_URL"))
}
