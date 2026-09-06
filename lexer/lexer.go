package lexer

import (
	"fmt"
	"os"

	"juni-interpreter/token"
)

func Lexer(filepath string) ([]token.Token, error) {
	source, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("read source file: %w", err)
	}

	return tokenize(source)
}
