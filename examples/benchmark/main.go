package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ClickHouse/ch-go"
	"github.com/ClickHouse/ch-go/proto"
)

func main() {
	// defer profile.Start().Stop()

	ctx := context.Background()

	conn, err := ch.Dial(ctx, ch.Options{})
	if err != nil {
		panic(err)
	}

	if err := conn.Do(ctx, ch.Query{
		Body: "DROP TABLE IF EXISTS benchmark2",
	}); err != nil {
		panic(err)
	}
	if err := conn.Do(ctx, ch.Query{
		Body: `CREATE TABLE benchmark2 (
			Col1 UInt64
		  , Col2 String
		  , Col3 Array(UInt8)
		  , Col4 DateTime
	  ) Engine Null`,
	}); err != nil {
		panic(err)
	}

	// Define all columns of table.
	var (
		col1 proto.ColUInt64
		col2 proto.ColStr
		col3 = new(proto.ColUInt8).Array()
		col4 = new(proto.ColDateTime)

		// now = time.Date(2010, 1, 1, 10, 22, 33, 345678, time.UTC)
	)

	start := time.Now()

	for i := 0; i < 1000_000; i++ {
		col1.Append(uint64(i))
		col2.Append("Golang SQL database driver")
		col3.Append([]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9})
		col4.Append(time.Now())
	}

	// Insert single data block.
	input := proto.Input{
		{Name: "Col1", Data: col1},
		{Name: "Col2", Data: col2},
		{Name: "Col3", Data: col3},
		{Name: "Col4", Data: col4},
	}
	if err := conn.Do(ctx, ch.Query{
		Body:  "INSERT INTO benchmark2 VALUES",
		Input: input,
	}); err != nil {
		panic(err)
	}

	fmt.Println(time.Since(start))

	// // Stream data to ClickHouse server in multiple data blocks.
	// var blocks int
	// if err := conn.Do(ctx, ch.Query{
	// 	Body:  input.Into("benchmark2"), // helper that generates INSERT INTO query with all columns
	// 	Input: input,

	// 	// OnInput is called to prepare Input data before encoding and sending
	// 	// to ClickHouse server.
	// 	OnInput: func(ctx context.Context) error {
	// 		// On OnInput call, you should fill the input data.
	// 		//
	// 		// NB: You should reset the input columns, they are
	// 		// not reset automatically.
	// 		//
	// 		// That is, we are re-using the same input columns and
	// 		// if we will return nil without doing anything, data will be
	// 		// just duplicated.

	// 		input.Reset() // calls "Reset" on each column

	// 		if blocks >= 10 {
	// 			// Stop streaming.
	// 			//
	// 			// This will also write tailing input data if any,
	// 			// but we just reset the input, so it is currently blank.
	// 			return io.EOF
	// 		}

	// 		// Append new values:
	// 		for i := 0; i < 10; i++ {
	// 			body.AppendBytes([]byte("Hello"))
	// 			ts.Append(now)
	// 			name.Append("name")
	// 			sevText.Append("DEBUG")
	// 			sevNumber.Append(10)
	// 			arr.Append([]string{"foo", "bar", "baz"})
	// 		}

	// 		// Data will be encoded and sent to ClickHouse server after returning nil.
	// 		// The Do method will return error if any.
	// 		blocks++
	// 		return nil
	// 	},
	// }); err != nil {
	// 	panic(err)
	// }
}
