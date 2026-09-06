package lexer

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"juni-interpreter/token"
)

func TestTokenize(t *testing.T) {
	testCases := []struct {
		name     string
		source   string
		expected []token.Token
		wantErr  bool
		eofLine  int
	}{
		{
			name:   "empty source",
			source: "",
		},
		{
			name:   "single-character tokens",
			source: "(){}[];,:.",
			expected: []token.Token{
				{Type: token.LEFT_PAREN, Lexeme: "(", Line: 1},
				{Type: token.RIGHT_PAREN, Lexeme: ")", Line: 1},
				{Type: token.LEFT_BRACE, Lexeme: "{", Line: 1},
				{Type: token.RIGHT_BRACE, Lexeme: "}", Line: 1},
				{Type: token.LEFT_BRACKET, Lexeme: "[", Line: 1},
				{Type: token.RIGHT_BRACKET, Lexeme: "]", Line: 1},
				{Type: token.SEMICOLON, Lexeme: ";", Line: 1},
				{Type: token.COMMA, Lexeme: ",", Line: 1},
				{Type: token.COLON, Lexeme: ":", Line: 1},
				{Type: token.DOT, Lexeme: ".", Line: 1},
			},
		},
		{
			name:   "arithmetic operators",
			source: "+-*/%",
			expected: []token.Token{
				{Type: token.PLUS, Lexeme: "+", Line: 1},
				{Type: token.MINUS, Lexeme: "-", Line: 1},
				{Type: token.STAR, Lexeme: "*", Line: 1},
				{Type: token.SLASH, Lexeme: "/", Line: 1},
				{Type: token.MODULO, Lexeme: "%", Line: 1},
			},
		},
		{
			name:   "comparison operators",
			source: "= == ! != > >= < <=",
			expected: []token.Token{
				{Type: token.EQUAL, Lexeme: "=", Line: 1},
				{Type: token.EQUAL_EQUAL, Lexeme: "==", Line: 1},
				{Type: token.BANG, Lexeme: "!", Line: 1},
				{Type: token.BANG_EQUAL, Lexeme: "!=", Line: 1},
				{Type: token.GREATER, Lexeme: ">", Line: 1},
				{Type: token.GREATER_EQUAL, Lexeme: ">=", Line: 1},
				{Type: token.LESS, Lexeme: "<", Line: 1},
				{Type: token.LESS_EQUAL, Lexeme: "<=", Line: 1},
			},
		},
		{
			name:   "assignment operators",
			source: "+= -= *= /= %=",
			expected: []token.Token{
				{Type: token.PLUS_EQUAL, Lexeme: "+=", Line: 1},
				{Type: token.MINUS_EQUAL, Lexeme: "-=", Line: 1},
				{Type: token.STAR_EQUAL, Lexeme: "*=", Line: 1},
				{Type: token.SLASH_EQUAL, Lexeme: "/=", Line: 1},
				{Type: token.MODULO_EQUAL, Lexeme: "%=", Line: 1},
			},
		},
		{
			name:   "integer number",
			source: "123",
			expected: []token.Token{
				{Type: token.NUMBER, Lexeme: "123", Line: 1},
			},
		},
		{
			name:   "zero",
			source: "0",
			expected: []token.Token{
				{Type: token.NUMBER, Lexeme: "0", Line: 1},
			},
		},
		{
			name:   "identifier",
			source: "abc",
			expected: []token.Token{
				{Type: token.IDENTIFIER, Lexeme: "abc", Line: 1},
			},
		},
		{
			name:   "identifier with underscore",
			source: "_abc",
			expected: []token.Token{
				{Type: token.IDENTIFIER, Lexeme: "_abc", Line: 1},
			},
		},
		{
			name:   "identifier with digits",
			source: "abc123",
			expected: []token.Token{
				{Type: token.IDENTIFIER, Lexeme: "abc123", Line: 1},
			},
		},
		{
			name:   "keywords",
			source: "if else for while function return var true false null",
			expected: []token.Token{
				{Type: token.IF, Lexeme: "if", Line: 1},
				{Type: token.ELSE, Lexeme: "else", Line: 1},
				{Type: token.FOR, Lexeme: "for", Line: 1},
				{Type: token.IDENTIFIER, Lexeme: "while", Line: 1},
				{Type: token.FUNCTION, Lexeme: "function", Line: 1},
				{Type: token.RETURN, Lexeme: "return", Line: 1},
				{Type: token.VAR, Lexeme: "var", Line: 1},
				{Type: token.TRUE, Lexeme: "true", Line: 1},
				{Type: token.FALSE, Lexeme: "false", Line: 1},
				{Type: token.NULL, Lexeme: "null", Line: 1},
			},
		},
		{
			name:   "string",
			source: `"hello"`,
			expected: []token.Token{
				{Type: token.STRING, Lexeme: `"hello"`, Literal: "hello", Line: 1},
			},
		},
		{
			name:   "string containing spaces",
			source: `"hello world"`,
			expected: []token.Token{
				{Type: token.STRING, Lexeme: `"hello world"`, Literal: "hello world", Line: 1},
			},
		},
		{
			name:   "empty string",
			source: `""`,
			expected: []token.Token{
				{Type: token.STRING, Lexeme: `""`, Literal: "", Line: 1},
			},
		},
		{
			name:    "multiline string",
			source:  "\"hello\nworld\"",
			eofLine: 2,
			expected: []token.Token{
				{Type: token.STRING, Lexeme: "\"hello\nworld\"", Literal: "hello\nworld", Line: 1},
			},
		},
		{
			name:   "variable declaration",
			source: "var x = 123;",
			expected: []token.Token{
				{Type: token.VAR, Lexeme: "var", Line: 1},
				{Type: token.IDENTIFIER, Lexeme: "x", Line: 1},
				{Type: token.EQUAL, Lexeme: "=", Line: 1},
				{Type: token.NUMBER, Lexeme: "123", Line: 1},
				{Type: token.SEMICOLON, Lexeme: ";", Line: 1},
			},
		},
		{
			name:    "tokens across lines",
			source:  "var\nx\n=\n10;",
			eofLine: 4,
			expected: []token.Token{
				{Type: token.VAR, Lexeme: "var", Line: 1},
				{Type: token.IDENTIFIER, Lexeme: "x", Line: 2},
				{Type: token.EQUAL, Lexeme: "=", Line: 3},
				{Type: token.NUMBER, Lexeme: "10", Line: 4},
				{Type: token.SEMICOLON, Lexeme: ";", Line: 4},
			},
		},
		{
			name:    "unterminated string",
			source:  `"unterminated`,
			wantErr: true,
		},
		{
			name:    "unexpected character",
			source:  "$",
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual, err := tokenize([]byte(testCase.source))
			if testCase.wantErr {
				if err == nil {
					t.Fatal("tokenize() returned no error, want an error")
				}
				return
			}

			if err != nil {
				t.Fatalf("tokenize() returned an unexpected error: %v", err)
			}

			eofLine := testCase.eofLine
			if eofLine == 0 {
				eofLine = 1
			}
			expected := append([]token.Token(nil), testCase.expected...)
			expected = append(expected, token.Token{Type: token.EOF, Line: eofLine})

			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("tokenize() = %#v, want %#v", actual, expected)
			}
		})
	}
}

func TestLexer(t *testing.T) {
	t.Run("reads and tokenizes a source file", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "program.juni")
		if err := os.WriteFile(path, []byte("var answer = 42;"), 0o600); err != nil {
			t.Fatalf("write temporary source file: %v", err)
		}

		actual, err := Lexer(path)
		if err != nil {
			t.Fatalf("Lexer() returned an unexpected error: %v", err)
		}

		expected := []token.Token{
			{Type: token.VAR, Lexeme: "var", Line: 1},
			{Type: token.IDENTIFIER, Lexeme: "answer", Line: 1},
			{Type: token.EQUAL, Lexeme: "=", Line: 1},
			{Type: token.NUMBER, Lexeme: "42", Line: 1},
			{Type: token.SEMICOLON, Lexeme: ";", Line: 1},
			{Type: token.EOF, Line: 1},
		}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Lexer() = %#v, want %#v", actual, expected)
		}
	})

	t.Run("returns an error for a missing source file", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "missing.juni")
		if _, err := Lexer(path); err == nil {
			t.Fatal("Lexer() returned no error for a missing source file")
		}
	})
}
