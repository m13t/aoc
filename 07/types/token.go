package types

type TokenType byte

const (
	TokenCommand TokenType = iota + 1
	TokenOutput
)

type Token struct {
	Type TokenType
	Args []string
}
