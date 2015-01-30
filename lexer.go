package ll

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type Lexer interface {
	NextToken() Token
}

const (
	EOF    = iota - 1
	NAME   = 2
	COMMA  = 3
	LBRACK = 4
	RBRACK = 5
	EQUAL  = 6
)

type ListLexer struct {
	input  string
	reader *bufio.Reader
}

func NewListLexer(input string) *ListLexer {
	return &ListLexer{
		input:  input,
		reader: bufio.NewReader(strings.NewReader(input)),
	}
}

func (l *ListLexer) NextToken() *Token {
	c, err := l.reader.ReadByte()

	for err == nil {
		switch string(c) {
		case " ", "\t", "\r", "\n":
		case ",":
			return &Token{COMMA, ","}
		case "[":
			return &Token{LBRACK, "["}
		case "]":
			return &Token{RBRACK, "]"}
		case "=":
			return &Token{EQUAL, "="}
		default:
			if isletter := unicode.IsLetter(rune(c)); isletter {
				names := []byte{}
				for unicode.IsLetter(rune(c)) {
					names = append(names, c)
					c, _ = l.reader.ReadByte()
				}

				token := &Token{NAME, string(names)}
				l.reader.UnreadByte()
				return token
			}

			errors.New(fmt.Sprintf("ll:invalid character %s", c))
		}

		c, err = l.reader.ReadByte()
	}

	return &Token{EOF, "EOF"}
}

type Token struct {
	typ  int
	text string
}
