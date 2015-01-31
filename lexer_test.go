package ll

import (
	"testing"
)

func TestListLexer(t *testing.T) {
	lexer := NewListLexer("[a,bb=c,[d,e]]")

	if token := lexer.NextToken(); token.text != "[" {
		t.Fatalf("expected [ actual is %s", token.text)
	}

	if token := lexer.NextToken(); token.text != "a" {
		t.Fatalf("expected a actual is %s", token.text)
	}

	if token := lexer.NextToken(); token.text != "," {
		t.Fatalf("expected , actual is %s", token.text)
	}
}
