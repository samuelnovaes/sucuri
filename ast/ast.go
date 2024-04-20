package ast

import "fmt"

const (
	NUMBER     = iota
	STRING     = iota
	BOOLEAN    = iota
	NULL       = iota
	IDENTIFIER = iota
	CALL       = iota
	FUNCTION   = iota
	RETURN     = iota
	LIB        = iota
)

type Expression interface {
	GetKind() int
	GetValue() any
}

type Null struct {
	Expression
}

func (node *Null) GetKind() int {
	return NULL
}

func (node *Null) GetValue() any {
	return nil
}

type Number struct {
	Expression
	Value float64
}

func (node *Number) GetKind() int {
	return NUMBER
}

func (node *Number) GetValue() any {
	return node.Value
}

type String struct {
	Expression
	Value string
}

func (node *String) GetKind() int {
	return STRING
}

func (node *String) GetValue() any {
	return node.Value
}

type Boolean struct {
	Expression
	Value bool
}

func (node *Boolean) GetKind() int {
	return BOOLEAN
}

func (node *Boolean) GetValue() any {
	return node.Value
}

type Identifier struct {
	Expression
	Symbol   string
	Evaluate bool
}

func (node *Identifier) GetKind() int {
	return IDENTIFIER
}

func (node *Identifier) GetValue() any {
	return fmt.Sprintf("[pointer %s]", node.Symbol)
}

type Lib struct {
	Expression
	Symbol string
}

func (node *Lib) GetKind() int {
	return LIB
}

func (node *Lib) GetValue() any {
	return fmt.Sprintf("[function %s]", node.Symbol)
}

type Call struct {
	Expression
	Args   []Expression
	Caller Expression
}

func (node *Call) GetKind() int {
	return CALL
}

func (node *Call) GetValue() any {
	return fmt.Sprintf("[call %v]", node.Caller.GetValue())
}

type Function struct {
	Expression
	Args []Identifier
	Body []Expression
}

func (node *Function) GetKind() int {
	return FUNCTION
}

func (node *Function) GetValue() any {
	return "[function]"
}

type Return struct {
	Expression
	Value Expression
}

func (node *Return) GetKind() int {
	return RETURN
}

func (node *Return) GetValue() any {
	return fmt.Sprintf("[return %v]", node.Value)
}
