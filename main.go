package main

import (
    "fmt"
    "os"

    "github.com/ohhfishal/gospace/lexer"
)

func main() {
    args := os.Args[1:]
    if len(args) != 1 {
        fmt.Println("usage: gospace FILE")
        return
    }
    l, err := lexer.NewLexerFromFile(args[0])
    if err != nil {
        fmt.Println(err)
        return
    }
    lexer.Dump(l)
}
