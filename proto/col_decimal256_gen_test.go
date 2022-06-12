// Code generated by ./cmd/ch-gen-col, DO NOT EDIT.

package proto

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ClickHouse/ch-go/internal/gold"
)

func TestColDecimal256_DecodeColumn(t *testing.T) {
	t.Parallel()
	const rows = 50
	var data ColDecimal256
	for i := 0; i < rows; i++ {
		v := Decimal256FromInt(i)
		data.Append(v)
		require.Equal(t, v, data.Row(i))
	}

	var buf Buffer
	data.EncodeColumn(&buf)
	t.Run("Golden", func(t *testing.T) {
		t.Parallel()
		gold.Bytes(t, buf.Buf, "col_decimal256")
	})
	t.Run("Ok", func(t *testing.T) {
		br := bytes.NewReader(buf.Buf)
		r := NewReader(br)

		var dec ColDecimal256
		require.NoError(t, dec.DecodeColumn(r, rows))
		require.Equal(t, data, dec)
		require.Equal(t, rows, dec.Rows())
		dec.Reset()
		require.Equal(t, 0, dec.Rows())
		require.Equal(t, ColumnTypeDecimal256, dec.Type())
	})
	t.Run("ZeroRows", func(t *testing.T) {
		r := NewReader(bytes.NewReader(nil))

		var dec ColDecimal256
		require.NoError(t, dec.DecodeColumn(r, 0))
	})
	t.Run("ErrUnexpectedEOF", func(t *testing.T) {
		r := NewReader(bytes.NewReader(nil))

		var dec ColDecimal256
		require.ErrorIs(t, dec.DecodeColumn(r, rows), io.ErrUnexpectedEOF)
	})
	t.Run("NoShortRead", func(t *testing.T) {
		var dec ColDecimal256
		requireNoShortRead(t, buf.Buf, colAware(&dec, rows))
	})
	t.Run("ZeroRowsEncode", func(t *testing.T) {
		var v ColDecimal256
		v.EncodeColumn(nil) // should be no-op
	})
}

func TestColDecimal256Array(t *testing.T) {
	const rows = 50
	data := NewArrDecimal256()
	for i := 0; i < rows; i++ {
		data.Append([]Decimal256{
			Decimal256FromInt(i),
			Decimal256FromInt(i + 1),
			Decimal256FromInt(i + 2),
		})
	}

	var buf Buffer
	data.EncodeColumn(&buf)
	t.Run("Golden", func(t *testing.T) {
		gold.Bytes(t, buf.Buf, "col_arr_decimal256")
	})
	t.Run("Ok", func(t *testing.T) {
		br := bytes.NewReader(buf.Buf)
		r := NewReader(br)

		dec := NewArrDecimal256()
		require.NoError(t, dec.DecodeColumn(r, rows))
		require.Equal(t, data, dec)
		require.Equal(t, rows, dec.Rows())
		dec.Reset()
		require.Equal(t, 0, dec.Rows())
		require.Equal(t, ColumnTypeDecimal256.Array(), dec.Type())
	})
	t.Run("ErrUnexpectedEOF", func(t *testing.T) {
		r := NewReader(bytes.NewReader(nil))

		dec := NewArrDecimal256()
		require.ErrorIs(t, dec.DecodeColumn(r, rows), io.ErrUnexpectedEOF)
	})
}

func BenchmarkColDecimal256_DecodeColumn(b *testing.B) {
	const rows = 1_000
	var data ColDecimal256
	for i := 0; i < rows; i++ {
		data = append(data, Decimal256FromInt(i))
	}

	var buf Buffer
	data.EncodeColumn(&buf)

	br := bytes.NewReader(buf.Buf)
	r := NewReader(br)

	var dec ColDecimal256
	if err := dec.DecodeColumn(r, rows); err != nil {
		b.Fatal(err)
	}
	b.SetBytes(int64(len(buf.Buf)))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		br.Reset(buf.Buf)
		r.raw.Reset(br)
		dec.Reset()

		if err := dec.DecodeColumn(r, rows); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkColDecimal256_EncodeColumn(b *testing.B) {
	const rows = 1_000
	var data ColDecimal256
	for i := 0; i < rows; i++ {
		data = append(data, Decimal256FromInt(i))
	}

	var buf Buffer
	data.EncodeColumn(&buf)

	b.SetBytes(int64(len(buf.Buf)))
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		buf.Reset()
		data.EncodeColumn(&buf)
	}
}
