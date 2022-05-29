// Code generated by ./cmd/ch-gen-col, DO NOT EDIT.

package proto

// ColDate represents Date column.
type ColDate []Date

// Compile-time assertions for ColDate.
var (
	_ ColInput  = ColDate{}
	_ ColResult = (*ColDate)(nil)
	_ Column    = (*ColDate)(nil)
)

// Type returns ColumnType of Date.
func (ColDate) Type() ColumnType {
	return ColumnTypeDate
}

// Rows returns count of rows in column.
func (c ColDate) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColDate) Row(i int) Date {
	return c[i]
}

// Append Date to column.
func (c *ColDate) Append(v Date) {
	*c = append(*c, v)
}

// AppendArr appends slice of Date to column.
func (c *ColDate) AppendArr(v []Date) {
	*c = append(*c, v...)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColDate) Reset() {
	*c = (*c)[:0]
}

// LowCardinality returns LowCardinality for Date .
func (c *ColDate) LowCardinality() *ColLowCardinalityOf[Date] {
	return &ColLowCardinalityOf[Date]{
		index: c,
	}
}

// Array is helper that creates Array of Date.
func (c *ColDate) Array() *ColArrOf[Date] {
	return &ColArrOf[Date]{
		Data: c,
	}
}

// NewArrDate returns new Array(Date).
func NewArrDate() *ColArr {
	return &ColArr{
		Data: new(ColDate),
	}
}

// AppendDate appends slice of Date to Array(Date).
func (c *ColArr) AppendDate(data []Date) {
	d := c.Data.(*ColDate)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}
