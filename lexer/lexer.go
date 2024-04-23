package lexer

import (
	"fmt"
	"strings"

	"github.com/samuelnovaes/sucuri/token"
)

const OPERATORS_ALPHABET = "!=+-*/%<>&|?$"

func shift(code *string) string {
	char := (*code)[0]
	*code = (*code)[1:]
	return string(char)
}

func isAlpha(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '_'
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func isAphaNum(char byte) bool {
	return isAlpha(char) || isDigit(char)
}

func isSkippable(char byte) bool {
	return char == ' ' || char == '\t' || char == '\r' || char == '\n'
}

func readNumber(code *string) token.Token {
	num := ""
	isFloat := false
	for len(*code) > 0 && (isDigit((*code)[0]) || (!isFloat && (*code)[0] == '.' && len(*code) > 1 && isDigit((*code)[1]))) {
		if (*code)[0] == '.' {
			isFloat = true
		}
		num += shift(code)
	}
	return token.Token{Literal: num, Type: token.NUMBER}
}

func readIdentifier(code *string) token.Token {
	ident := ""
	for len(*code) > 0 && (isAphaNum((*code)[0]) || (len(ident) > 0 && ident[len(ident)-1] != '.' && (*code)[0] == '.' && len(*code) > 1 && isAlpha((*code)[1]))) {
		ident += shift(code)
	}
	var tokenType int
	switch ident {
	case "true":
		tokenType = token.BOOLEAN
	case "false":
		tokenType = token.BOOLEAN
	case "null":
		tokenType = token.NULL
	default:
		tokenType = token.IDENTIFIER
	}
	return token.Token{Literal: ident, Type: tokenType}
}

func readString(code *string) token.Token {
	tk := token.Token{Literal: "", Type: token.STRING}
	quote := (*code)[0]
	tk.Literal += shift(code)
	for len(*code) > 0 {
		if (*code)[0] == quote {
			quotedStr := tk.Literal + string(quote)
			for strings.HasSuffix(quotedStr, "\\\\"+string(quote)) {
				quotedStr = strings.ReplaceAll(quotedStr, "\\\\", "")
			}
			if !strings.HasSuffix(quotedStr, "\\"+string(quote)) {
				tk.Literal += shift(code)
				return tk
			}
		}
		tk.Literal += shift(code)
	}
	return tk
}

func readPointerValue(code *string) token.Token {
	shift(code)
	tk := readIdentifier(code)
	tk.Type = token.POINTER_VALUE
	return tk
}

func readNegativeNumber(code *string) token.Token {
	num := ""
	num += shift(code)
	num += readNumber(code).Literal
	return token.Token{Type: token.NUMBER, Literal: num}
}

func readCall(code *string, tokens *[]token.Token) {
	*tokens = append(*tokens, token.Token{Literal: shift(code), Type: token.OPEN_PAREM})
	ident := ""
	for isSkippable((*code)[0]) {
		shift(code)
	}
	for strings.ContainsRune(OPERATORS_ALPHABET, rune((*code)[0])) {
		ident += shift(code)
	}
	if ident != "" {
		*tokens = append(*tokens, token.Token{Literal: ident, Type: token.IDENTIFIER})
	}
}

func Tokenize(code string) []token.Token {
	tokens := []token.Token{}
	for len(code) > 0 {
		if code[0] == '(' {
			readCall(&code, &tokens)
		} else if code[0] == ')' {
			tokens = append(tokens, token.Token{Literal: shift(&code), Type: token.CLOSE_PAREM})
		} else if code[0] == '{' {
			tokens = append(tokens, token.Token{Literal: shift(&code), Type: token.OPEN_BRACE})
		} else if code[0] == '}' {
			tokens = append(tokens, token.Token{Literal: shift(&code), Type: token.CLOSE_BRACE})
		} else if isDigit(code[0]) {
			tokens = append(tokens, readNumber(&code))
		} else if isAlpha(code[0]) {
			tokens = append(tokens, readIdentifier(&code))
		} else if code[0] == '"' {
			tokens = append(tokens, readString(&code))
		} else if code[0] == '$' {
			tokens = append(tokens, readPointerValue(&code))
		} else if code[0] == '-' {
			tokens = append(tokens, readNegativeNumber(&code))
		} else if isSkippable(code[0]) {
			shift(&code)
		} else {
			panic(fmt.Sprintf("Unexpected token: %c", code[0]))
		}
	}
	return tokens
}
