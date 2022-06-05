package redidash_test

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/rbtyang/godash/redidash"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

var _redis *redis.Client

func init() {
	log.Println("Before this tests")
	_redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "foobared", // no password set
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
