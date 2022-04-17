// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

// ColEnum16 represents Enum16 column.
type ColEnum16 []Enum16

// Compile-time assertions for ColEnum16.
var (
	_ ColInput  = ColEnum16{}
	_ ColResult = (*ColEnum16)(nil)
	_ Column    = (*ColEnum16)(nil)
)

// Type returns ColumnType of Enum16.
func (ColEnum16) Type() ColumnType {
	return ColumnTypeEnum16
}

// Rows returns count of rows in column.
func (c ColEnum16) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColEnum16) Row(i int) Enum16 {
	return c[i]
}

// Append Enum16 to column.
func (c *ColEnum16) Append(v Enum16) {
	*c = append(*c, v)
}

// AppendArr appends slice of Enum16 to column.
func (c *ColEnum16) AppendArr(v []Enum16) {
	*c = append(*c, v...)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColEnum16) Reset() {
	*c = (*c)[:0]
}

// LowCardinality returns LowCardinality for Enum16 .
func (c *ColEnum16) LowCardinality() *ColLowCardinalityOf[Enum16] {
	return &ColLowCardinalityOf[Enum16]{
		index: c,
	}
}

// NewArrEnum16 returns new Array(Enum16).
func NewArrEnum16() *ColArr {
	return &ColArr{
		Data: new(ColEnum16),
	}
}

// AppendEnum16 appends slice of Enum16 to Array(Enum16).
func (c *ColArr) AppendEnum16(data []Enum16) {
	d := c.Data.(*ColEnum16)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}
