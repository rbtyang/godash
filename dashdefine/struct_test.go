package dashdefine_test

import (
	"github.com/rbtyang/godash/dashdefine"
	"github.com/rbtyang/godash/dashjson"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

/*
@Editor robotyang at 2024

TestEmpty is a ...
*/
func TestEmpty(t *testing.T) {
	st := dashdefine.Empty{}
	{
		ty := reflect.TypeOf(st)
		kd := ty.Kind()
		assert.Equal(t, reflect.Struct, kd)
	}
	{
		js, err := dashjson.Marshal(st)
		assert.Equal(t, nil, err)
		assert.Equal(t, "{}", js)
	}
}
