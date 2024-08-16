package parser

import (
    "fmt"
    "io"

    "github.com/ohhfishal/gospace/lexer"
    "github.com/ohhfishal/gospace/lexer/token"
    "github.com/ohhfishal/gospace/parser/ir"
)

func Parse(scanner lexer.Lexer) (*ir.Program, error) {
    return NewParser(scanner).Parse()
}

func NewParser(scanner lexer.Lexer) *Parser {
    return &Parser{lexer: scanner}
}

type Parser struct {
    lexer lexer.Lexer
    program *ir.Program
    err error
}

type ParseError struct {
    Message string
}
func (self ParseError) Error() string {
    return self.Message
}

func (self *Parser) consume(matchType token.Type) *token.Token {
    next, err := self.lexer.Pop()
    if err != nil {
        panic(err)
    }
    if next.Type != matchType {
        err = ParseError{fmt.Sprintf("Found: [%s] at %d.\nExpected: [%s].", next.Type, next.Position, matchType)}
        panic(err)
    }
    return next
}

func (self *Parser) peek() token.Type {
    next, err := self.lexer.Peek()
    if err != nil {
        panic(err)
    }
    return next.Type
}


func (self *Parser) Parse() (program *ir.Program, err error) {
    defer func() {
        caught := recover()
        if caught  == io.EOF {
            program = self.program
            err = nil
            return
        }
        err = caught.(error)
    }()

    for {
        self.consume(token.WHITE_SPACE)
    }

}
