package lib

import (
	"github.com/samuelnovaes/sucuri/ast"
	"github.com/samuelnovaes/sucuri/lib/arithmetic"
	"github.com/samuelnovaes/sucuri/lib/assignment"
	"github.com/samuelnovaes/sucuri/lib/comparison"
	"github.com/samuelnovaes/sucuri/lib/logical"
	"github.com/samuelnovaes/sucuri/lib/pointer"
	"github.com/samuelnovaes/sucuri/lib/std"
	"github.com/samuelnovaes/sucuri/lib/structural"
)

var Lib = map[string]ast.LibFunction{
	"print":   std.Print,
	"println": std.Println,
	"format":  std.Format,
	"input":   std.Input,
	"number":  std.Number,
	"string":  std.String,
	"return":  structural.Return,
	"if":      structural.If,
	"elif":    structural.Elif,
	"else":    structural.Else,
	"var":     structural.Var,
	"const":   structural.Const,
	"while":   structural.While,
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
	"$":       pointer.EvalPtr,
}
