package ast

import (
	"fmt"
)

const (
	NUMBER     = iota
	STRING     = iota
	BOOLEAN    = iota
	NULL       = iota
	IDENTIFIER = iota
	CALL       = iota
	FUNCTION   = iota
	LIB        = iota
	RETURN     = iota
	REFERENCE  = iota
	BREAK      = iota
	CONTINUE   = iota
)

type LibFunction func(ctx *Context, args ...Expression) Expression

type Context map[string]*Reference

type Expression interface {
	GetKind() int
	GetValue(*Context) any
	ToString(ctx *Context) string
}

// REFERENCE ===========================================================
//

type Reference struct {
	Expression
	Value Expression
	Const bool
}

func (node *Reference) GetKind() int {
	return REFERENCE
}

func (node *Reference) GetValue(ctx *Context) any {
	return node.Value
}

func (node *Reference) ToString(ctx *Context) string {
	return fmt.Sprintf("%p", &node.Value)
}

//
// NULL ================================================================
//

type Null struct {
	Expression
}

func (node *Null) GetKind() int {
	return NULL
}

func (node *Null) GetValue(ctx *Context) any {
	return nil
}

func (node *Null) ToString(ctx *Context) string {
	return "null"
}

//
// NUMBER ==============================================================
//

type Number struct {
	Expression
	Value float64
}

func (node *Number) GetKind() int {
	return NUMBER
}

func (node *Number) GetValue(ctx *Context) any {
	return node.Value
}

func (node *Number) ToString(ctx *Context) string {
	return fmt.Sprintf("%v", node.Value)
}

//
// STRING ==============================================================
//

type String struct {
	Expression
	Value string
}

func (node *String) GetKind() int {
	return STRING
}

func (node *String) GetValue(ctx *Context) any {
	return node.Value
}

func (node *String) ToString(ctx *Context) string {
	return node.Value
}

//
// BOOLEAN =============================================================
//

type Boolean struct {
	Expression
	Value bool
}

func (node *Boolean) GetKind() int {
	return BOOLEAN
}

func (node *Boolean) GetValue(ctx *Context) any {
	return node.Value
}

func (node *Boolean) ToString(ctx *Context) string {
	if node.Value {
		return "true"
	}
	return "false"
}

//
// IDENTIFIER ==========================================================
//

type Identifier struct {
	Expression
	Symbol   string
	Evaluate bool
}

func (node *Identifier) GetKind() int {
	return IDENTIFIER
}

func (node *Identifier) GetValue(ctx *Context) any {
	ref, ok := (*ctx)[node.Symbol]
	if ok {
		return ref.Value
	}
	return Null{}
}

func (node *Identifier) ToString(ctx *Context) string {
	ref, ok := (*ctx)[node.Symbol]
	if ok {
		return fmt.Sprintf("%p", &ref.Value)
	}
	return (&Null{}).ToString(ctx)
}

//
// CALL ================================================================
//

type Call struct {
	Expression
	Args   []Expression
	Caller Expression
}

func (node *Call) GetKind() int {
	return CALL
}

func (node *Call) GetValue(ctx *Context) any {
	return node
}

func (node *Call) ToString(ctx *Context) string {
	return fmt.Sprintf("%p", &node)
}

//
// FUNCTION ============================================================
//

type Function struct {
	Expression
	Args []Identifier
	Body []Expression
}

func (node *Function) GetKind() int {
	return FUNCTION
}

func (node *Function) GetValue(ctx *Context) any {
	return node
}

func (node *Function) ToString(ctx *Context) string {
	return fmt.Sprintf("%p", &node)
}

//
// LIB =================================================================
//

type Lib struct {
	Expression
	Function *LibFunction
}

func (node *Lib) GetKind() int {
	return LIB
}

func (node *Lib) GetValue(ctx *Context) any {
	return node.Function
}

func (node *Lib) ToString(ctx *Context) string {
	return fmt.Sprintf("%p", &node)
}

//
// RETURN ==============================================================
//

type Return struct {
	Expression
	Value Expression
}

func (node *Return) GetKind() int {
	return RETURN
}

func (node *Return) GetValue(ctx *Context) any {
	return node
}

func (node *Return) ToString(ctx *Context) string {
	return fmt.Sprintf("%p", &node)
}

//
// BREAK ===============================================================
//

type Break struct {
	Expression
}

func (node *Break) GetKind() int {
	return BREAK
}

func (node *Break) GetValue(ctx *Context) any {
	return node
}

func (node *Break) ToString(ctx *Context) string {
	return fmt.Sprintf("%p", &node)
}

//
// CONTINUE ============================================================
//

type Continue struct {
	Expression
}

func (node *Continue) GetKind() int {
	return CONTINUE
}

func (node *Continue) GetValue(ctx *Context) any {
	return node
}

func (node *Continue) ToString(ctx *Context) string {
	return fmt.Sprintf("%p", &node)
}

//
// =====================================================================

func GetRef(value Expression, ctx *Context) *Reference {
	if value.GetKind() == REFERENCE {
		return value.(*Reference)
	}
	symbol := value.(*Identifier).Symbol
	if ref, ok := (*ctx)[symbol]; ok {
		return ref
	}
	return nil
}
