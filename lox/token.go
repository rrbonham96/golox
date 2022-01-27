package lox

import "fmt"

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

func NewToken(t_type TokenType, lexeme string, literal interface{}, line int) Token {
	return Token{
		Type:    t_type,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    line,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("%s %s %s", t.Type, t.Lexeme, t.Literal)
}
