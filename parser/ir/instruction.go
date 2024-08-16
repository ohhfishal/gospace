package ir

type Type uint16
const (
    OP_PUSH = iota
    OP_DUPLICATE
    OP_COPY
    OP_SWAP
    OP_DISCARD
    OP_SLIDE
    OP_ADD
    OP_SUBTRACT
    OP_MULTIPLY
    OP_INT_DIVIDE
    OP_MODULO
    OP_STORE
    OP_RETRIEVE
    OP_MARK
    OP_CALL
    OP_JUMP
    OP_JUMP_IF_ZERO
    OP_JUMP_IF_NEGATIVE
    OP_RETURN
    OP_END
    OP_PRINT_CHAR
    OP_PRINT_NUM
    OP_READ_CHAR
    OP_READ_NUM
)

type Label string
type Number string

type Instruction struct {
    Type Type
    Label *Label
    Number *Number
}

type Program struct {
    Instructions []Instruction
}
