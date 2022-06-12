// Code generated by ./cmd/ch-gen-col, DO NOT EDIT.

package proto

// ColEnum8 represents Enum8 column.
type ColEnum8 []Enum8

// Compile-time assertions for ColEnum8.
var (
	_ ColInput  = ColEnum8{}
	_ ColResult = (*ColEnum8)(nil)
	_ Column    = (*ColEnum8)(nil)
)

// Type returns ColumnType of Enum8.
func (ColEnum8) Type() ColumnType {
	return ColumnTypeEnum8
}

// Rows returns count of rows in column.
func (c ColEnum8) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColEnum8) Row(i int) Enum8 {
	return c[i]
}

// Append Enum8 to column.
func (c *ColEnum8) Append(v Enum8) {
	*c = append(*c, v)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColEnum8) Reset() {
	*c = (*c)[:0]
}

// LowCardinality returns LowCardinality for Enum8 .
func (c *ColEnum8) LowCardinality() *ColLowCardinalityOf[Enum8] {
	return &ColLowCardinalityOf[Enum8]{
		index: c,
	}
}

// Array is helper that creates Array of Enum8.
func (c *ColEnum8) Array() *ColArrOf[Enum8] {
	return &ColArrOf[Enum8]{
		Data: c,
	}
}

// NewArrEnum8 returns new Array(Enum8).
func NewArrEnum8() *ColArrOf[Enum8] {
	return &ColArrOf[Enum8]{
		Data: new(ColEnum8),
	}
}
