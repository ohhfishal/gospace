package lexer

import (
    "bufio"
    "fmt"
    "os"
    "io"

    "github.com/ohhfishal/gospace/lexer/token"
)

type Lexer interface {
    Pop() (*token.Token, error)
    Peek() (*token.Token, error)
}

func ReadAll(self Lexer) []*token.Token {
    tokens := []*token.Token{}
    cur, err := self.Pop()
    for err == nil {
        tokens = append(tokens, cur)
        cur, err = self.Pop()
    }

    if err != io.EOF {
        panic(err)
    }
    return tokens
}

func Dump(self Lexer) []*token.Token {
    tokens := ReadAll(self)
    fmt.Println(len(tokens))
    for i, t := range tokens {
        fmt.Printf("[%d] %s\n", i, t.String())
    }
    return tokens

}

func NewLexerFromFile(filename string) (*ReaderLexer, error) {
    reader, err := os.Open(filename)
    bReader := bufio.NewReader(reader)
    if err != nil || bReader == nil {
        return nil, err
    }
    return &ReaderLexer{reader: bReader}, nil
}

type ReaderLexer struct {
    reader *bufio.Reader
    tokenCache *token.Token
    bytesRead int
}

func (self *ReaderLexer) Pop() (*token.Token, error) {
    return self.nextToken(true)
}

func (self *ReaderLexer) Peek() (*token.Token, error) {
    return self.nextToken(false)

}

func (self *ReaderLexer) nextToken(consume bool) (next *token.Token, err error) {
    if self.tokenCache != nil {
        next = self.tokenCache
    } else {
        char, err := self.nextRune()
        if err != nil {
            return nil, err
        }

        next = &token.Token{
            Type: 1 << char,
            Position: self.bytesRead,
        }
    }

    if consume {
        self.tokenCache = nil
    } else {
        self.tokenCache = next
    }
    return next, err
}

func (self *ReaderLexer) nextRune() (rune, error) {
    char, err := self.readRune()
    for (err == nil && !token.IsValid(char)) {
        char, err = self.readRune()
    }
    return char, err
}

func (self *ReaderLexer) readRune() (rune, error) {
    char, count, err := self.reader.ReadRune()
    self.bytesRead = self.bytesRead + count
    return char, err
}
