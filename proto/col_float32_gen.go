// Code generated by ./cmd/ch-gen-col, DO NOT EDIT.

package proto

// ColFloat32 represents Float32 column.
type ColFloat32 []float32

// Compile-time assertions for ColFloat32.
var (
	_ ColInput  = ColFloat32{}
	_ ColResult = (*ColFloat32)(nil)
	_ Column    = (*ColFloat32)(nil)
)

// Rows returns count of rows in column.
func (c ColFloat32) Rows() int {
	return len(c)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColFloat32) Reset() {
	*c = (*c)[:0]
}

// Type returns ColumnType of Float32.
func (ColFloat32) Type() ColumnType {
	return ColumnTypeFloat32
}

// Row returns i-th row of column.
func (c ColFloat32) Row(i int) float32 {
	return c[i]
}

// Append float32 to column.
func (c *ColFloat32) Append(v float32) {
	*c = append(*c, v)
}

// LowCardinality returns LowCardinality for Float32 .
func (c *ColFloat32) LowCardinality() *ColLowCardinality[float32] {
	return &ColLowCardinality[float32]{
		index: c,
	}
}

// Array is helper that creates Array of float32.
func (c *ColFloat32) Array() *ColArr[float32] {
	return &ColArr[float32]{
		Data: c,
	}
}

// Nullable is helper that creates Nullable(float32).
func (c *ColFloat32) Nullable() *ColNullable[float32] {
	return &ColNullable[float32]{
		Values: c,
	}
}

// NewArrFloat32 returns new Array(Float32).
func NewArrFloat32() *ColArr[float32] {
	return &ColArr[float32]{
		Data: new(ColFloat32),
	}
}
