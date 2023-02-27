package astdash_test

import (
	"github.com/rbtyang/godash/astdash"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

/*
TestParseFile is a ...

@Editor robotyang at 2023
*/
func TestParseFile(t *testing.T) {
	{
		a := astdash.NewAst()
		err := a.ParseFile("../arrdash/arr.go")
		if err != nil {
			log.Fatal(err)
		}
		assert.Equal(t, "arrdash", a.Package)
		assert.Contains(t, a.Imports[0].Name, "logdash")
		assert.Equal(t, "Contain", a.Funcs[1].Name)
	}
}
