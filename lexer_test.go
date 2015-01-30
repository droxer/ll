package ll

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {

	lexer := NewListLexer("[a,bb=c,[d,e]]")
	fmt.Println(lexer.NextToken())
	fmt.Println(lexer.NextToken())
	fmt.Println(lexer.NextToken())
	fmt.Println(lexer.NextToken())
	fmt.Println(lexer.NextToken())
	fmt.Println(lexer.NextToken())
	fmt.Println(lexer.NextToken())
	fmt.Println(lexer.NextToken())
	fmt.Println(lexer.NextToken())
	fmt.Println(lexer.NextToken())
	fmt.Println(lexer.NextToken())
	fmt.Println(lexer.NextToken())
	fmt.Println(lexer.NextToken())
	fmt.Println(lexer.NextToken())
}
