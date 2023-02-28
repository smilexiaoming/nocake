package converter

import (
	"fmt"
)

func main() {
	t2t := NewTable2Struct()
	t2t.TagKey("gorm")
	t2t.Prefix("t_")
	err := t2t.
		// SavePath("./model.go").
		Dsn("root:Zwx_1994@tcp(localhost:3306)/nocake?charset=utf8").
		Run()
	fmt.Println(err)
}
