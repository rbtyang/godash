package dashast_test

import (
	"fmt"
	"github.com/rbtyang/godash/dashast"
	"log"
)

/*
ExampleNewAst is a ...
*/
func ExampleNewAst() {
	a := dashast.NewAst()
	err := a.ParseFile("../dasharr/arr.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a.Package)
	fmt.Println(a.Imports[0].Name)
	fmt.Println(a.Funcs[0].Name)
	//Output:
	//dasharr
	//dashlog
	//Contain
}
