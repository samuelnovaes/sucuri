package lib

import (
	"github.com/samuelnovaes/sucuri/ast"
)

var Lib = map[string]ast.LibFunction{
	"set":     Set,
	"print":   Print,
	"println": Println,
	"format":  Format,
	"return":  Return,
	"input":   Input,
	"number":  Number,
	"string":  String,
	"add":     Add,
	"sub":     Sub,
	"mul":     Mul,
	"div":     Div,
	"inc":     Inc,
	"dec":     Dec,
	"mod":     Mod,
	"neg":     Neg,
	"eq":      Eq,
	"ne":      Ne,
	"gt":      Gt,
	"lt":      Lt,
	"gte":     Gte,
	"lte":     Lte,
	"not":     Not,
	"if":      If,
	"elif":    Elif,
	"else":    Else,
}
