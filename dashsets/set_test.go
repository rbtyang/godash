package dashsets_test

import (
	"github.com/rbtyang/godash/dashsets"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
TestIntSet is a ...

@Editor robotyang at 2023
*/
func TestIntSet(t *testing.T) {
	var leng int
	var bools bool
	{
		set := dashsets.NewSet[int]() //这里要声明 T的实际类型 是int
		leng = set.Len()
		assert.Equal(t, 0, leng)

		set.Add(3)
		leng = set.Len()
		assert.Equal(t, 1, leng)

		set.Add(5)
		leng = set.Len()
		assert.Equal(t, 2, leng)

		set.Add(5)
		leng = set.Len()
		assert.Equal(t, 2, leng)

		bools = set.Contains(5)
		assert.Equal(t, true, bools)

		bools = set.Contains(9)
		assert.NotEqual(t, true, bools)

		{
			set2 := dashsets.NewSetWith[int](3, 5)
			bools = set.Equals(set2)
			assert.Equal(t, true, bools)
		}

		set.Remove(3)
		assert.Equal(t, 1, set.Len())
		{
			set2 := dashsets.NewSetWith[int](5)
			bools = set.Equals(set2)
			assert.Equal(t, true, bools)
		}
		{
			set2 := dashsets.NewSetWith[int](3)
			bools = set.Equals(set2)
			assert.NotEqual(t, true, bools)
		}

		set.Clear()
		leng = set.Len()
		assert.Equal(t, 0, leng)
		{
			set2 := dashsets.NewSetWith[int]()
			bools = set.Equals(set2)
			assert.Equal(t, true, bools)
		}
	}
}
