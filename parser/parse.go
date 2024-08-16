package parser

import (
    "fmt"

    "github.com/ohhfishal/gospace/lexer"
    "github.com/ohhfishal/gospace/lexer/token"
    "github.com/ohhfishal/gospace/parser/ir"
)

const WHITE_SPACE = token.WHITE_SPACE
const LINE_FEED = token.LINE_FEED
const TAB = token.TAB

func Parse(scanner lexer.Lexer) (*ir.Program, error) {
    return NewParser(scanner).Parse()
}

func NewParser(scanner lexer.Lexer) *Parser {
    var program ir.Program
    return &Parser{
        lexer: scanner,
        program: &program,
    }
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

func panicParseError(msg string) {
    // TODO: Make this a better message (Like add in more info)
    panic(ParseError{msg})
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
        if caught := recover(); caught != nil {
            err = caught.(ParseError)
        }
    }()
    
    for self.peek() != token.EOF {
        self.imp()
    }

    self.consume(token.EOF)
    return self.program, nil
}


func (self *Parser) imp() {
    switch self.peek() {
        case WHITE_SPACE:
            self.consume(WHITE_SPACE)
            self.stack()
        case LINE_FEED:
            self.consume(LINE_FEED)
            self.flow()
        case TAB:
            self.consume(TAB)
            self.tab()
        default:
            panicParseError("Expected IMP, got EOF")
    }
}

func (self *Parser) stack() {
    switch self.peek() {
        case WHITE_SPACE:
            self.consume(WHITE_SPACE)
            _ = self.number()
            // OP_PUSH
        case LINE_FEED:
            self.consume(LINE_FEED)
            switch self.peek() {
                case WHITE_SPACE:
                    self.consume(WHITE_SPACE)
                    // OP_DUPLICATE
                case LINE_FEED:
                    self.consume(LINE_FEED)
                    // OP_DISCARD
                case TAB:
                    self.consume(TAB)
                    // OP_SWAP
                default:
                    panicParseError("Stack-LF, got EOF")
            }
        case TAB:
            self.consume(TAB)
            switch self.peek() {
                case WHITE_SPACE:
                    self.consume(WHITE_SPACE)
                    _ = self.number()
                    // OP_COPY_NTH
                case LINE_FEED:
                    self.consume(LINE_FEED)
                    _ = self.number()
                    // OP_SLIDE
                default:
                    panicParseError("Stack-TAB, got EOF")
            }
        default:
            panicParseError("Stack, got EOF")
    }
}

func (self *Parser) flow() {
    // SS
    // ST
    // SL
    // TS
    // TT
    // TL
    // LL
}

func (self *Parser) tab() {
    // S = Math
    // T = Heap
    // L = IO
}

func (self *Parser) math() {
    // [Space][Space]  -   Addition
    // [Space][Tab]    -   Subtraction
    // [Space][LF] -   Multiplication
    // [Tab][Space]    -   Integer Division
    // [Tab][Tab]  -   Modulo
}

func (self *Parser) heap() {
    // S = store
    // Tab = retrieve
}

func (self *Parser) io_parse() {
    // [Space][Space]  -   Output the character at the top of the stack
    // [Space][Tab]    -   Output the number at the top of the stack
    // [Tab][Space]    -   Read a character and place it in the location given by the top of the stack
    // [Tab][Tab]  -   Read a number and place it in the location given by the top of the stack
}

func (self *Parser) number() *ir.Number {
// Numbers can be any number of bits wide, and are simply represented as a series of [Space] and [Tab], terminated by a [LF]. [Space] represents the binary digit 0, [Tab] represents 1. The sign of a number is given by its first character, [Space] for positive and [Tab] for negative. Note that this is not twos complement, it just indicates a sign. 
    return nil
}
