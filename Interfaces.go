package lexorank

//ILexoNumeralSystem Interface Design
type ILexoNumeralSystem interface {
	Name() string

	GetBase() int

	GetPositiveChar() byte

	GetNegativeChar() byte

	GetRadixPointChar() byte

	ToDigit(var1 byte) int

	ToChar(var1 int) byte
}
