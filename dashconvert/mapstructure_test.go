package dashconvert_test

import (
	"github.com/rbtyang/godash/dashconvert"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
WeakMapToStcWithTag is a ...

@Editor robotyang at 2023
*/
func TestWeakMapToStcWithTag(t *testing.T) {
	inMap := map[string]interface{}{
		"myName":   123,                      // number => string
		"myAge":    "42",                     // string => number
		"myEmails": map[string]interface{}{}, // empty map => empty array
		"xxName":   456,                      // number => string
		"xxAge":    "53",                     // string => number
		"xxEmails": map[string]interface{}{}, // empty map => empty array
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
		err := dashconvert.WeakMapToStcWithTag(inMap, &recv, "myTag")
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
		err := dashconvert.WeakMapToStcWithTag(inMap, &recv, "xxTag")
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, want, recv)
	}
}
