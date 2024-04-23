package lib

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/lib/arithmetic"
	"github.com/samuelnovaes/sucuri/lib/assignment"
	"github.com/samuelnovaes/sucuri/lib/comparison"
	"github.com/samuelnovaes/sucuri/lib/logical"
	"github.com/samuelnovaes/sucuri/lib/structural"
)

var Lib = map[string]ast.LibFunction{
	"print":   structural.Print,
	"println": structural.Println,
	"format":  structural.Format,
	"return":  structural.Return,
	"input":   structural.Input,
	"number":  structural.Number,
	"string":  structural.String,
	"if":      structural.If,
	"elif":    structural.Elif,
	"else":    structural.Else,
	"var":     structural.Var,
	"const":   structural.Const,
	"$":       structural.EvalPtr,
	"?":       structural.Ternary,
	"=":       assignment.Set,
	"++":      assignment.Inc,
	"--":      assignment.Dec,
	"+=":      assignment.SetAdd,
	"-=":      assignment.SetSub,
	"*=":      assignment.SetMul,
	"/=":      assignment.SetDiv,
	"//=":     assignment.SetIntDiv,
	"%=":      assignment.SetMod,
	"**=":     assignment.SetPow,
	"&&=":     assignment.SetAnd,
	"||=":     assignment.SetOr,
	"+":       arithmetic.Add,
	"-":       arithmetic.Sub,
	"*":       arithmetic.Mul,
	"/":       arithmetic.Div,
	"//":      arithmetic.IntDiv,
	"%":       arithmetic.Mod,
	"**":      arithmetic.Pow,
	"==":      comparison.Eq,
	"!=":      comparison.Ne,
	">":       comparison.Gt,
	"<":       comparison.Lt,
	">=":      comparison.Gte,
	"<=":      comparison.Lte,
	"!":       logical.Not,
	"&&":      logical.And,
	"||":      logical.Or,
}
