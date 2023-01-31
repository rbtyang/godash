package astdash_test

import (
	"fmt"
	"github.com/rbtyang/godash/astdash"
	"log"
)

func ExampleNewAst() {
	a := astdash.NewAst()
	err := a.ParseFile("../arrdash/arr.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a.Package)         //"arrdash"
	fmt.Println(a.Imports[0].Name) //"fmt"
	fmt.Println(a.Funcs[1].Name)   //"Contains"
}
