package ll

import (
	"testing"
)

func TestListParser(t *testing.T) {
	lexer := NewListLexer("[a,b=c,[d,e]]")
	parser := NewListParser(lexer, 2)
	parser.Parse()
}
