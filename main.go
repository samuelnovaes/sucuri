package main

import (
	"os"

	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/evaluator"
	"github.com/samuelnovaes/sucuri/lexer"
	"github.com/samuelnovaes/sucuri/lib"
	"github.com/samuelnovaes/sucuri/parser"
)

func main() {
	code, err := os.ReadFile("example.suc")
	if err != nil {
		panic(err.Error())
	}

	ctx := ast.Context{
		"#IF": nil,
	}
	for key, fn := range lib.Lib {
		ctx[key] = &ast.Lib{Symbol: key, Function: &fn}
	}

	tokens := lexer.Tokenize(string(code))
	program := parser.ProduceAST(tokens)
	evaluator.EvalFunction(program, &ctx)
}
