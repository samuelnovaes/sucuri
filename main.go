package main

import (
	"os"

	"github.com/samuelnovaes/sucuri/evaluator"
	"github.com/samuelnovaes/sucuri/lexer"
	"github.com/samuelnovaes/sucuri/parser"
)

func main() {
	code, err := os.ReadFile("example.suc")
	if err != nil {
		panic(err.Error())
	}
	tokens := lexer.Tokenize(string(code))
	program := parser.ProduceAST(tokens)
	evaluator.Eval(program)
}
