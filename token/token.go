package token

const (
	NUMBER      = iota
	STRING      = iota
	BOOLEAN     = iota
	NULL        = iota
	IDENTIFIER  = iota
	OPEN_PAREM  = iota
	CLOSE_PAREM = iota
	OPEN_BRACE  = iota
	CLOSE_BRACE = iota
	DOLLAR      = iota
)

type Token struct {
	Literal string
	Type    int
}
