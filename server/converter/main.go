package main

import (
	"fmt"

	"github.com/gohouse/converter"
)

func main() {
	t2t := converter.NewTable2Struct()

	err := t2t.
		SavePath("./model.go").
		Dsn("root:Zwx_1994@tcp(localhost:3306)/nocake?charset=utf8").
		Run()
	fmt.Println(err)
}
