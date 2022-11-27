package redidash_test

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/rbtyang/godash/filedash"
	"github.com/rbtyang/godash/redidash"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"testing"
	"time"
)

var _redis *redis.Client

const _configYaml = "temp/config.yaml"

type _configYamlSt struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}

var _configYamlObj = _configYamlSt{}

func init() {
	log.Println("Before lock_test.go tests")

	var err error
	var yamlFile *os.File

	if !filedash.IsExistFile(_configYaml) {
		yamlFile, err = filedash.Rebuild(_configYaml)
		if err != nil {
			log.Panicln(err)
			return
		}
		defer yamlFile.Close()
		if _, err := yamlFile.WriteString("host: localhost\nport: 6379\npassword: "); err != nil {
			log.Panicln(err)
			return
		}
	} else {
		yamlFile, err = os.OpenFile(_configYaml, os.O_RDONLY, os.ModePerm)
		if err != nil {
			log.Panicln(err)
			return
		}
		defer yamlFile.Close()
	}

	yamlFileByt, err := filedash.ReadByFile(yamlFile)
	if err != nil {
		log.Panicln(err)
		return
	}
	err = yaml.Unmarshal(yamlFileByt, &_configYamlObj)
	if err != nil {
		log.Panicln(err)
		return
	}

	_redis = redis.NewClient(&redis.Options{
		Addr:     _configYamlObj.Host + ":" + _configYamlObj.Port,
		Password: _configYamlObj.Password,
	})
}

func TestLock(t *testing.T) {
	{
		ctx := context.Background()
		lock := redidash.NewLock(ctx, _redis)
		key := "qwe123123"

		{ //first lock
			ok, err := lock.Lock(key, 55*time.Second)
			if err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, ok, true) //上锁成功
		}

		{ //conflict lock
			ok, err := lock.Lock(key, 55*time.Second)
			if err != nil {
				t.Error(err)
				return
			}
			//if !ok { //上锁失败（说明已经有锁）
			//	t.Error("请求频繁，请稍后再试")
			//	return
			//}
			assert.Equal(t, ok, false) //上锁冲突
		}

		{ //free lock
			_, err := lock.Free(key) //释放锁
			if err != nil {
				t.Error(err)
				return
			}
		}

		{ //again lock
			ok, err := lock.Lock(key, 55*time.Second)
			if err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, ok, true) //再次上锁成功
		}

		{ //again free lock
			_, err := lock.Free(key) //再次释放锁
			if err != nil {
				t.Error(err)
				return
			}
		}

	}
}
