package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/token"
)

func shift(tokens *[]token.Token) token.Token {
	token := (*tokens)[0]
	*tokens = (*tokens)[1:]
	return token
}

func parseNumber(tokens *[]token.Token) ast.Expression {
	literal := shift(tokens).Literal
	value, error := strconv.ParseFloat(literal, 64)
	if error != nil {
		panic(fmt.Sprintf("Can not convert '%s' to float64", literal))
	}
	return &ast.Number{Value: value}
}

func parseString(tokens *[]token.Token) ast.Expression {
	literal := shift(tokens).Literal
	literal = strings.ReplaceAll(literal, "\n", "\\n")
	value, error := strconv.Unquote(literal)
	if error != nil {
		panic(fmt.Sprintf("Can not convert '%s' to string", literal))
	}
	return &ast.String{Value: value}
}

func parseBoolean(tokens *[]token.Token) ast.Expression {
	literal := shift(tokens).Literal
	if literal == "true" {
		return &ast.Boolean{Value: true}
	}
	return &ast.Boolean{Value: false}
}

func parseIdentifier(tokens *[]token.Token, evaluate bool) ast.Expression {
	literal := shift(tokens).Literal
	return &ast.Identifier{Symbol: literal, Evaluate: evaluate}
}

func parseBraceBody(tokens *[]token.Token) []ast.Expression {
	body := []ast.Expression{}
	shift(tokens)
	for len(*tokens) > 0 && (*tokens)[0].Type != token.CLOSE_BRACE {
		body = append(body, parse(tokens))
	}
	shift(tokens)
	return body
}

func parseFunction(tokens *[]token.Token) ast.Expression {
	shift(tokens)
	body := []ast.Expression{}
	args := []ast.Identifier{}
	for len(*tokens) > 0 && (*tokens)[0].Type != token.CLOSE_BRACKET {
		args = append(args, *parse(tokens).(*ast.Identifier))
	}
	shift(tokens)
	if len(*tokens) > 0 {
		if (*tokens)[0].Type == token.OPEN_BRACE {
			body = parseBraceBody(tokens)
		} else {
			body = []ast.Expression{&ast.Call{
				Args:   []ast.Expression{parse(tokens)},
				Caller: &ast.Identifier{Symbol: "return", Evaluate: false},
			}}
		}
	}
	return &ast.Function{Args: args, Body: body}
}

func parseColonFunction(tokens *[]token.Token) ast.Expression {
	shift(tokens)
	body := []ast.Expression{&ast.Call{
		Args:   []ast.Expression{parse(tokens)},
		Caller: &ast.Identifier{Symbol: "return", Evaluate: false},
	}}
	return &ast.Function{Args: []ast.Identifier{}, Body: body}
}

func parseBraceFunction(tokens *[]token.Token) ast.Expression {
	return &ast.Function{Args: []ast.Identifier{}, Body: parseBraceBody(tokens)}
}

func parseCall(tokens *[]token.Token) ast.Expression {
	shift(tokens)
	args := []ast.Expression{}
	var caller ast.Expression
	for len(*tokens) > 0 && (*tokens)[0].Type != token.CLOSE_PAREM {
		if caller == nil {
			caller = parse(tokens)
		} else {
			args = append(args, parse(tokens))
		}
	}
	shift(tokens)
	return &ast.Call{Caller: caller, Args: args}
}

func parse(tokens *[]token.Token) ast.Expression {
	currentToken := (*tokens)[0]
	switch currentToken.Type {
	case token.NUMBER:
		return parseNumber(tokens)
	case token.STRING:
		return parseString(tokens)
	case token.BOOLEAN:
		return parseBoolean(tokens)
	case token.IDENTIFIER:
		return parseIdentifier(tokens, false)
	case token.POINTER_VALUE:
		return parseIdentifier(tokens, true)
	case token.OPEN_PAREM:
		return parseCall(tokens)
	case token.OPEN_BRACKET:
		return parseFunction(tokens)
	case token.COLON:
		return parseColonFunction(tokens)
	case token.OPEN_BRACE:
		return parseBraceFunction(tokens)
	case token.NULL:
		expr := &ast.Null{}
		shift(tokens)
		return expr
	default:
		panic(fmt.Sprintf("Unexpected token: %s", currentToken.Literal))
	}
}

func ProduceAST(tokens []token.Token) ast.Function {
	body := []ast.Expression{}
	for len(tokens) > 0 {
		body = append(body, parse(&tokens))
	}
	return ast.Function{Args: []ast.Identifier{}, Body: body}
}
