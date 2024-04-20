package lib

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/context"
)

type Function func(*context.Context, ...ast.Expression) ast.Expression

var Lib = map[string]Function{
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
}
