package astdash_test

import (
	"github.com/rbtyang/godash/astdash"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestParseFile(t *testing.T) {
	{
		a := astdash.NewAst()
		err := a.ParseFile("../arrdash/arr.go")
		if err != nil {
			log.Fatal(err)
		}
		assert.Equal(t, "arrdash", a.Package)
		assert.Equal(t, "fmt", a.Imports[0].Name)
		assert.Equal(t, "Contains", a.Funcs[1].Name)
	}
}
