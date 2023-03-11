package dashast_test

import (
	"github.com/rbtyang/godash/dashast"
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
		a := dashast.NewAst()
		err := a.ParseFile("../dasharr/arr.go")
		if err != nil {
			log.Fatal(err)
		}
		assert.Equal(t, "dasharr", a.Package)
		assert.Contains(t, a.Imports[0].Name, "dashlog")
		assert.Equal(t, "Contain", a.Funcs[1].Name)
	}
}
