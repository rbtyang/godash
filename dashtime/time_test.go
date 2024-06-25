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
func TestLayout(t *testing.T) {
	//自定义格式
	{
		want := dashtime.Lay_YmDHMS
		recv, err := dashtime.Layout("2116-11-12 15:14:15")
		assert.Equal(t, nil, err)
		assert.Equal(t, want, recv)
	}
	{
		want := dashtime.Lay_YmDHM
		recv, err := dashtime.Layout("2116-11-12 15:14")
		assert.Equal(t, nil, err)
		assert.Equal(t, want, recv)
	}
	{
		want := dashtime.Lay_YmD
		recv, err := dashtime.Layout("2116-11-12")
		assert.Equal(t, nil, err)
		assert.Equal(t, want, recv)
	}

	{
		want := dashtime.Lay_YmDHMS_Zh
		recv, err := dashtime.Layout("2116年11月12日15:14:15")
		assert.Equal(t, nil, err)
		assert.Equal(t, want, recv)
	}
	{
		want := dashtime.Lay_YmDHM_Zh
		recv, err := dashtime.Layout("2116年11月12日15:14")
		assert.Equal(t, nil, err)
		assert.Equal(t, want, recv)
	}
	{
		want := dashtime.Lay_YmD_Zh
		recv, err := dashtime.Layout("2116年11月12日")
		assert.Equal(t, nil, err)
		assert.Equal(t, want, recv)
	}

	//内置标准格式
	{
		want := dashtime.Lay_RubyDate
		recv, err := dashtime.Layout("Wed Aug 12 15:14:15 -0711 2116")
		assert.Equal(t, nil, err)
		assert.Equal(t, want, recv)
	}
	{
		want := dashtime.Lay_RFC1123
		recv, err := dashtime.Layout("Wed, 12 Aug 2116 15:14:15 MST")
		assert.Equal(t, nil, err)
		assert.Equal(t, want, recv)
	}

	//错误情况
	{
		want := ""
		recv, err := dashtime.Layout("11月12日")
		assert.EqualError(t, err, "not yet supported layout")
		assert.Equal(t, want, recv)
	}
}

/*
@Editor robotyang at 2024

TestParse is a ...
*/
func TestParse(t *testing.T) {
	want := getTestTime()
	{
		recv, err := dashtime.Parse("2024-05-01")
		assert.Equal(t, nil, err)
		assert.Equal(t, want, recv) //TODO debug
	}
	{
		recv, err := dashtime.Parse("2024-05-01 12:34:56")
		assert.Equal(t, nil, err)
		assert.Equal(t, want, recv)
	}
	{
		recv, err := dashtime.Parse("2024-05-01 12:34:56")
		assert.Equal(t, nil, err)
		assert.Equal(t, want, recv)
	}
	{
		recv, err := dashtime.Parse("2024-05-01 12:34:56")
		assert.Equal(t, nil, err)
		assert.Equal(t, want, recv)
	}
}

/*
@Editor robotyang at 2024

TestParse is a ...
*/
func TestParseLoc(t *testing.T) {
	{
		want := dashtime.Loc_Local
		recv, err := dashtime.ParseLoc("2024-05-01", dashtime.Loc_Local)
		assert.Equal(t, nil, err)
		assert.Equal(t, "2024-05-01 00:00:00 +0800 CST", recv.String())
		assert.Equal(t, want, recv.Location().String())
	}
	{
		want := dashtime.Loc_China
		recv, err := dashtime.ParseLoc("2024-05-01", dashtime.Loc_China)
		assert.Equal(t, nil, err)
		assert.Equal(t, "2024-05-01 00:00:00 +0800 CST", recv.String())
		assert.Equal(t, want, recv.Location().String())
	}
	{
		want := dashtime.Loc_China
		recv, err := dashtime.ParseLoc("2024-05-01 12:34", dashtime.Loc_China)
		assert.Equal(t, nil, err)
		assert.Equal(t, "2024-05-01 12:34:00 +0800 CST", recv.String())
		assert.Equal(t, want, recv.Location().String())
	}
	{
		want := dashtime.Loc_China
		recv, err := dashtime.ParseLoc("2024-05-01 12:34:56", dashtime.Loc_China)
		assert.Equal(t, nil, err)
		assert.Equal(t, "2024-05-01 12:34:56 +0800 CST", recv.String())
		assert.Equal(t, want, recv.Location().String())
	}
	{
		want := dashtime.Loc_America
		recv, err := dashtime.ParseLoc("2024-05-01 12:34:56", dashtime.Loc_America)
		assert.Equal(t, nil, err)
		assert.Equal(t, "2024-05-01 12:34:56 -0700 PDT", recv.String())
		assert.Equal(t, want, recv.Location().String())
	}
}

/*
@Editor robotyang at 2024

Format is a ...
*/
func TestDuraNextDawn(t *testing.T) {
	nowDateStr := time.Now().Format(dashtime.Lay_Date)
	{
		tim, err := dashtime.Parse(nowDateStr + " 13:15:30")
		dura := dashtime.DuraNextDawn(tim)
		dH := dura.Hours()
		dM := dura.Minutes()
		dS := dura.Seconds()
		assert.Equal(t, nil, err)
		assert.Equal(t, 10.741666666666667, dH)
		assert.Equal(t, 644.5, dM)
		assert.Equal(t, 38670, dS)
	}
	{
		tim, err := dashtime.Parse(nowDateStr + " 01:02:03")
		dura := dashtime.DuraNextDawn(tim)
		dH := dura.Hours()
		dM := dura.Minutes()
		dS := dura.Seconds()
		assert.Equal(t, nil, err)
		assert.Equal(t, 22.965833333333332, dH)
		assert.Equal(t, 1377.95, dM)
		assert.Equal(t, 82677, dS)
	}
}
