// token/token.go
package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILEGAL = "ILEGAL"
	EOF    = "EOF"

	// 識別子 + リテラル
	IDENT = "IDENT" // add, foobar, x, y,
	INT   = "INT"   // 123456

	// 演算子
	ASSIGN = "="
	PLUS   = "+"

	// デリミタ
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
