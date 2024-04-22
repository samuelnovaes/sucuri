package ast

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
)

type LibFunction func(*Context, ...Expression) Expression
type Context map[string]Expression

type Expression interface {
	GetKind() int
	GetValue(*Context) any
}

type Null struct {
	Expression
}

func (node *Null) GetKind() int {
	return NULL
}

func (node *Null) GetValue(ctx *Context) any {
	return nil
}

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

type Identifier struct {
	Expression
	Symbol   string
	Evaluate bool
}

func (node *Identifier) GetKind() int {
	return IDENTIFIER
}

func (node *Identifier) GetValue(ctx *Context) any {
	expr := (*ctx)[node.Symbol]
	return &expr
}

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

type Lib struct {
	Expression
	Symbol   string
	Function *LibFunction
}

func (node *Lib) GetKind() int {
	return LIB
}

func (node *Lib) GetValue(ctx *Context) any {
	return node.Function
}

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
