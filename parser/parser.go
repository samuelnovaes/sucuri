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

func parseFunction(tokens *[]token.Token, caller ast.Expression, args []ast.Expression) ast.Expression {
	funcArgs := []ast.Identifier{}
	if caller != nil {
		funcArgs = append(funcArgs, *caller.(*ast.Identifier))
	}
	for _, arg := range args {
		funcArgs = append(funcArgs, *arg.(*ast.Identifier))
	}
	body := []ast.Expression{}
	shift(tokens)
	for len(*tokens) > 0 && (*tokens)[0].Type != token.CLOSE_BRACE {
		body = append(body, parse(tokens))
	}
	shift(tokens)
	return &ast.Function{Args: funcArgs, Body: body}
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
	if len(*tokens) > 0 && (*tokens)[0].Type == token.OPEN_BRACE {
		return parseFunction(tokens, caller, args)
	}
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
