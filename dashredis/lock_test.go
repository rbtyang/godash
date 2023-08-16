package dashredis_test

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/rbtyang/godash/dashfile"
	"github.com/rbtyang/godash/dashredis"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"testing"
	"time"
)

var redisCli *redis.Client

type configSt struct {
	Redis *redisSt `yaml:"redis"` //这里Redis首字母要大写，否则读不到配置
}
type redisSt struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

var config = configSt{&redisSt{}}

/*
init is a ...
*/
func init() {
	log.Println("Before lock_test.go tests")

	var err error
	var yamlFile *os.File

	yamlFile, err = os.OpenFile("../test.yaml", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Panicln(err)
		return
	}
	defer yamlFile.Close()

	yamlFileByt, err := dashfile.ReadByFile(yamlFile)
	if err != nil {
		log.Panicln(err)
		return
	}
	err = yaml.Unmarshal(yamlFileByt, &config)
	if err != nil {
		log.Panicln(err)
		return
	}

	redisCli = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host + ":" + config.Redis.Port,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})
}

/*
TestLock is a ...
*/
func TestLock(t *testing.T) {
	{
		ctx := context.Background()
		lock := dashredis.NewLock(ctx, redisCli)
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
