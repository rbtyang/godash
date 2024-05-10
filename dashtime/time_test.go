package dashtime_test

import (
	"github.com/rbtyang/godash/dashtime"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func getTestTime() *time.Time {
	loc, err := time.LoadLocation(dashtime.Loc_China)
	if err != nil {
		panic(err)
	}
	t1 := time.Date(2024, 5, 1, 12, 34, 56, 0, loc)
	return &t1
}

/*
@Editor robotyang at 2024

TestParseToChina is a ...
*/
func TestParseToChina(t *testing.T) {
	want := getTestTime()
	recv, err := dashtime.ParseToChina("2024-05-01 12:34:56")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, want, recv)
}

/*
@Editor robotyang at 2024

Format is a ...
*/
func TestFormat(t *testing.T) {
	t1 := getTestTime()
	{
		want := "2024-05-01"
		recv := dashtime.Format(t1, dashtime.Fmt_Date)
		assert.Equal(t, want, recv)
	}
	{
		want := "12:34:56"
		recv := dashtime.Format(t1, dashtime.Fmt_Time)
		assert.Equal(t, want, recv)
	}
	{
		want := "2024-05-01 12:34:56"
		recv := dashtime.Format(t1)
		assert.Equal(t, want, recv)
	}
	{
		want := "2024年05月01日12:34:56"
		recv := dashtime.Format(t1, dashtime.Fmt_Zh_Y_m_D_H_M_S)
		assert.Equal(t, want, recv)
	}
}
