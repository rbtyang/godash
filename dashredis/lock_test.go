package dashredis_test

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/rbtyang/godash/dashredis"
	"github.com/rbtyang/godash/internal"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

var redisCli *redis.Client

/*
init is a ...
*/
func init() {
	config, err := internal.LoadConfig()
	if err != nil {
		log.Panicln(err)
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
			res, err := lock.Unlock(key) //释放锁
			if err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, res, int64(1))
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
			res, err := lock.UnlockScript(key) //再次释放锁
			if err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, res, int64(1))
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
			res, err := lock.Unlock(key) //再次释放锁(值存在)
			if err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, res, int64(1))
		}

		{ //again free lock
			res, err := lock.Unlock(key) //再再次释放锁(值不存在时)
			if err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, res, int64(0))
		}

		{ //again free lock
			res, err := lock.UnlockScript(key) //再再再次释放锁(值不存在时)
			if err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, res, int64(0))
		}

	}
}
