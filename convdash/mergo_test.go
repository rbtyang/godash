package convdash

import (
	"fmt"
	"github.com/imdario/mergo"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Config struct {
	DbName   string
	DbUser   string
	DbPwd    string
	MaxConn  int32
	MinConn  int32
	IdleConn int32
	Label    []string
}

func TestMergo(t *testing.T) {
	{
		dstConfig := Config{
			DbName:   "111sfsffas",
			DbUser:   "111listsi",
			DbPwd:    "111zs123fasd",
			MaxConn:  11110,
			MinConn:  1117,
			IdleConn: 1119,
			Label:    []string{"111asdfa", "11112334", "111asdf1334"},
		}
		srcConfig := Config{
			DbName:   "222sfsffas",
			DbUser:   "222zhangsan",
			DbPwd:    "", //没值
			MaxConn:  0, //没值
			MinConn:  2227,
			IdleConn: 2229,
			Label:    []string{"222asdfa", "2222334", "222asdf1334"},
		}
		//src有值则覆盖dst，src没值则保留dst
		if err := mergo.Merge(&dstConfig, srcConfig, mergo.WithOverride); err != nil {
			panic(fmt.Errorf("Apollo MergeConfig error: %s \n", err))
		}

		want := Config{
			DbName:   "222sfsffas",
			DbUser:   "222zhangsan",
			DbPwd:    "111zs123fasd",
			MaxConn:  11110,
			MinConn:  2227,
			IdleConn: 2229,
			Label:    []string{"222asdfa", "2222334", "222asdf1334"},
		}
		recv := dstConfig
		assert.Equal(t, want, recv)
	}
}
