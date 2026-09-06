package lexer

import (
	"fmt"
	"juni-interpreter/token"
)

func tokenize(source []byte) ([]token.Token, error) {
	code := string(source)
	start, current := 0, 0
	tokens := []token.Token{}
	var atLine int = 1
	for current < len(code) {
		{
			nextChar, err := peek(current, code)
			if err != nil {
				fmt.Printf("On Last Character: %v\n", err)
			}
			switch code[current] {
			case ' ', '\r', '\t':
				advance(&current, &code)
			case '\n':
				atLine++
				advance(&current, &code)
			case '[':
				addToken("[", token.LEFT_BRACKET, atLine, &tokens)
				advance(&current, &code)

			case ']':
				addToken("]", token.RIGHT_BRACKET, atLine, &tokens)
				advance(&current, &code)

			case '{':
				addToken("{", token.LEFT_BRACE, atLine, &tokens)
				advance(&current, &code)

			case '}':
				addToken("}", token.RIGHT_BRACE, atLine, &tokens)
				advance(&current, &code)

			case ';':
				addToken(";", token.SEMICOLON, atLine, &tokens)
				advance(&current, &code)

			case '(':
				addToken("(", token.LEFT_PAREN, atLine, &tokens)
				advance(&current, &code)
			case ')':
				addToken(")", token.RIGHT_PAREN, atLine, &tokens)
				advance(&current, &code)

			case ',':
				addToken(",", token.COMMA, atLine, &tokens)
				advance(&current, &code)
			case ':':
				addToken(":", token.COLON, atLine, &tokens)
				advance(&current, &code)

			case '.':
				addToken(".", token.DOT, atLine, &tokens)
				advance(&current, &code)

			case '+':
				if nextChar == '=' {
					addToken("+=", token.PLUS_EQUAL, atLine, &tokens)
					advance(&current, &code)
					advance(&current, &code)
				} else {
					addToken("+", token.PLUS, atLine, &tokens)
					advance(&current, &code)
				}
			case '-':
				if nextChar == '=' {
					addToken("-=", token.MINUS_EQUAL, atLine, &tokens)
					advance(&current, &code)
					advance(&current, &code)
				} else {
					addToken("-", token.MINUS, atLine, &tokens)
					advance(&current, &code)
				}

			case '*':
				if nextChar == '=' {
					addToken("*=", token.STAR_EQUAL, atLine, &tokens)
					advance(&current, &code)
					advance(&current, &code)

				} else {
					advance(&current, &code)
					addToken("*", token.STAR, atLine, &tokens)
				}
			case '/':
				if nextChar == '=' {
					addToken("/=", token.SLASH_EQUAL, atLine, &tokens)
					advance(&current, &code)
					advance(&current, &code)
				} else {
					advance(&current, &code)
					addToken("/", token.SLASH, atLine, &tokens)
				}
			case '%':
				if nextChar == '=' {
					addToken("%=", token.MODULO_EQUAL, atLine, &tokens)
					advance(&current, &code)
					advance(&current, &code)
				} else {
					addToken("%", token.MODULO, atLine, &tokens)
					advance(&current, &code)
				}
			case '=':
				if nextChar == '=' {
					advance(&current, &code)
					advance(&current, &code)
					addToken("==", token.EQUAL_EQUAL, atLine, &tokens)
				} else {
					advance(&current, &code)
					addToken("=", token.EQUAL, atLine, &tokens)
				}
			case '!':
				if nextChar == '=' {
					advance(&current, &code)
					advance(&current, &code)
					addToken("!=", token.BANG_EQUAL, atLine, &tokens)
				} else {
					advance(&current, &code)
					addToken("!", token.BANG, atLine, &tokens)
				}
			case '>':
				if nextChar == '=' {
					advance(&current, &code)
					advance(&current, &code)
					addToken(">=", token.GREATER_EQUAL, atLine, &tokens)
				} else {
					advance(&current, &code)
					addToken(">", token.GREATER, atLine, &tokens)
				}
			case '<':
				if nextChar == '=' {
					advance(&current, &code)
					advance(&current, &code)
					addToken("<=", token.LESS_EQUAL, atLine, &tokens)
				} else {
					advance(&current, &code)
					addToken("<", token.LESS, atLine, &tokens)
				}
			case '"':
				advance(&current, &code)
				for current < len(code) && code[current] != '"' {
					advance(&current, &code)
				}
				if current >= len(code) {
					return nil, fmt.Errorf("unterminated string literal")
				}
				advance(&current, &code)
				addToken(code[start:current], token.STRING, atLine, &tokens)

			default:
				if code[current] >= '0' && code[current] <= '9' {
					for current < len(code) && (code[current] >= '0' && code[current] <= '9') {
						advance(&current, &code)
					}
					addToken(code[start:current], token.NUMBER, atLine, &tokens)
				} else if (code[current] >= 'a' && code[current] <= 'z') || (code[current] >= 'A' && code[current] <= 'Z') || code[current] == '_' {
					for current < len(code) && ((code[current] >= 'a' && code[current] <= 'z') || (code[current] >= 'A' && code[current] <= 'Z') || (code[current] >= '0' && code[current] <= '9') || code[current] == '_') {
						advance(&current, &code)
					}

					addToken(code[start:current], token.IDENTIFIER, atLine, &tokens)
				} else {
					fmt.Printf("unexpected character")
					advance(&current, &code)
				}
			}
			start = current
		}

	}

	return tokens, nil
}
