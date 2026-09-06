package lexer

import (
	"fmt"

	"juni-interpreter/token"
)

func peek(idx int, code string) (byte, error) {
	if idx+1 >= len(code) {
		return 0, fmt.Errorf("peek out of bounds")
	}
	return code[idx+1], nil
}

func advance(idx *int, code *string) error {
	if *idx >= len(*code) {
		return fmt.Errorf("advance out of bounds")
	}
	*idx++
	return nil
}

func addToken(lexeme string, tokenType token.TokenType, atLine int, tokens *[]token.Token) {
	if tokenType == token.IDENTIFIER {
		keywordMap := map[string]token.TokenType{
			"if":       token.IF,
			"else":     token.ELSE,
			"for":      token.FOR,
			"function": token.FUNCTION,
			"return":   token.RETURN,
			"var":      token.VAR,
			"true":     token.TRUE,
			"false":    token.FALSE,
			"null":     token.NULL,
		}

		if keywordType, isKeyword := keywordMap[lexeme]; isKeyword {
			tokenType = keywordType
		}
	}

	var literal any
	if tokenType == token.STRING {
		literal = lexeme[1 : len(lexeme)-1]
	}

	t := token.Token{
		Type:    tokenType,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    atLine,
	}
	*tokens = append(*tokens, t)
}
