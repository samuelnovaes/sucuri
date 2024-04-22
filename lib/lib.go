package lib

import (
	"github.com/samuelnovaes/sucuri/ast"
)

var Lib = map[string]ast.LibFunction{
	"print":   Print,
	"println": Println,
	"format":  Format,
	"return":  Return,
	"input":   Input,
	"number":  Number,
	"string":  String,
	"if":      If,
	"elif":    Elif,
	"else":    Else,
	"=":       Set,
	"+":       Add,
	"-":       Sub,
	"*":       Mul,
	"/":       Div,
	"++":      Inc,
	"--":      Dec,
	"%":       Mod,
	"**":      Pow,
	"==":      Eq,
	"!=":      Ne,
	">":       Gt,
	"<":       Lt,
	">=":      Gte,
	"<=":      Lte,
	"!":       Not,
}
