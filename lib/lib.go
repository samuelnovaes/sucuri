package lib

import (
	"github.com/samuelnovaes/sucuri/ast"
)

var Lib = map[string]ast.LibFunction{
	"set":    Set,
	"print":  Print,
	"format": Format,
	"return": Return,
	"input":  Input,
	"number": Number,
	"string": String,
	"add":    Add,
	"sub":    Sub,
	"mul":    Mul,
	"div":    Div,
	"inc":    Inc,
	"dec":    Dec,
	"mod":    Mod,
	"pow":    Pow,
	"neg":    Neg,
	"eq":     Eq,
}
