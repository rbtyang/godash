package dashenv_test

import (
	"github.com/rbtyang/godash/dashenv"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

/*
@Editor robotyang at 2023

TestEnv is a ...
*/
func TestEnv(t *testing.T) {
	keyArr := []string{"GO_ENV", "RUN_ENV"}
	for _, key := range keyArr {
		err := os.Setenv(key, dashenv.Test)
		if err != nil {
			t.Fatal(err)
		}
		dashenv.Init(key)

		{
			want := false
			recv := dashenv.IsDev()
			assert.Equal(t, want, recv)
		}
		{
			want := true
			recv := dashenv.IsTest()
			assert.Equal(t, want, recv)
		}
	}
}
