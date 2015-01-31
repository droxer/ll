package ll

import (
	"errors"
	"fmt"
)

type Parser interface {
	Parse()
}

type ListParser struct {
	lexer  Lexer
	tokens []*Token
	k      int
	p      int
}

func NewListParser(lexer Lexer, k int) *ListParser {
	parser := &ListParser{
		lexer:  lexer,
		tokens: make([]*Token, k),
		k:      k,
		p:      0,
	}

	parser.init()

	return parser
}

func (l *ListParser) Parse() {
	l.list()
}

func (l *ListParser) list() {
	l.match(LBRACK)
	l.elements()
	l.match(RBRACK)
}

func (l *ListParser) match(x int) {
	if l.lookahead(1).typ != x {
		errors.New(fmt.Sprintf("unexpecting found %s \n", l.lookahead(1)))
	}

	l.consume()
}

func (l *ListParser) elements() {
	l.element()
	for l.lookahead(1).typ == COMMA {
		l.match(COMMA)
		l.element()
	}
}

func (l *ListParser) element() {
	switch {
	case l.lookahead(1).typ == NAME && l.lookahead(2).typ == EQUAL:
		l.match(NAME)
		l.match(EQUAL)
		l.match(NAME)
	case l.lookahead(1).typ == NAME:
		l.match(NAME)
	case l.lookahead(1).typ == LBRACK:
		l.list()
	default:
		errors.New(fmt.Sprintf("expecting name or list; found %s \n", l.lookahead(1).text))
	}
}

func (l *ListParser) init() {
	for i := 1; i <= l.k; i++ {
		l.consume()
	}
}

func (l *ListParser) lookahead(i int) *Token {
	return l.tokens[(l.p+i-1)%l.k]
}

func (l *ListParser) consume() {
	l.tokens[l.p] = l.lexer.NextToken()
	l.p = (l.p + 1) % l.k
}
