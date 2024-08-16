package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args[1:]
    fmt.Println("Hello world")
    if len(args) != 1 {
        fmt.Println("usage: gospace FILE")
        return
    }
    fmt.Println(args)
}
