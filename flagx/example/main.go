package main

import (
	"fmt"

	"github.com/fengde/gocommon/flagx"
	"github.com/fengde/gocommon/jsonx"
)

func main() {
	var input struct {
		Age   int     `flag:"age" default:"1" help:"年龄"`
		User  string  `flag:"user" default:"fedel"  help:"用户名称"`
		Money float64 `flag:"money" help:"金钱"`
		Old   bool    `flag:"old" help:"是不是老人家"`
	}

	err := flagx.Parse(&input)

	fmt.Println(err, jsonx.MarshalToStringNoErr(input))
}
