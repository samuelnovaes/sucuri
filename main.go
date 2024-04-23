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
	code, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err.Error())
	}

	ctx := ast.Context{
		"#IF": nil,
	}
	for key, fn := range lib.Lib {
		lib := ast.Lib{Function: &fn}
		ctx[key] = &ast.Reference{Value: &lib, Const: true}
	}

	tokens := lexer.Tokenize(string(code))
	program := parser.ProduceAST(tokens)
	evaluator.EvalFunction(program, &ctx)
}
