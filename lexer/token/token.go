package token

type Type uint64
const (
    WHITE_SPACE = 1 << ' '
    TAB = 1 << '\t'
    LINE_FEED = 1 << '\n'
    EOF = 0
)

func IsValid(char rune) bool {
    return char == ' ' || char == '\n' || char == '\t'
}

type Token struct {
    Type Type
    Position int
}


func (self Token) String() string {
    return self.Type.String()
}
func (self Type) String() string {
    switch self {
        case WHITE_SPACE:
            return "SPACE"
        case TAB:
            return "TAB"
        case LINE_FEED:
            return "NEW LINE"
        default:
            return "INVALID"
    }
}
