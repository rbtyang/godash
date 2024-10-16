package dashconv_test

import (
	"testing"

	"github.com/rbtyang/godash/dashconv"
	"github.com/stretchr/testify/assert"
)

/*
WeakMapToStructWithTag is a ...
*/
func TestWeakMapToStcWithTag(t *testing.T) {
	inMap := map[string]any{
		"myName":   123,              // number => string
		"myAge":    "42",             // string => number
		"myEmails": map[string]any{}, // empty map => empty array
		"xxName":   456,              // number => string
		"xxAge":    "53",             // string => number
		"xxEmails": map[string]any{}, // empty map => empty array
	}
	type Person struct {
		Name   string   `json:"name" yaml:"name" mapstructure:"name" myTag:"myName" xxTag:"xxName"`
		Age    int      `json:"age" yaml:"age" mapstructure:"age" myTag:"myAge" xxTag:"xxAge"`
		Emails []string `json:"emails" yaml:"emails" mapstructure:"emails" myTag:"myEmails" xxTag:"xxEmails"`
	}

	{
		want := Person{
			Name:   "123",
			Age:    42,
			Emails: []string{},
		}
		recv := Person{}
		err := dashconv.WeakMapToStructWithTag(inMap, &recv, "myTag")
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, want, recv)
	}

	{
		want := Person{
			Name:   "456",
			Age:    53,
			Emails: []string{},
		}
		recv := Person{}
		err := dashconv.WeakMapToStructWithTag(inMap, &recv, "xxTag")
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, want, recv)
	}
}
