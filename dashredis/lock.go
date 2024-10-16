package dashredis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type lock struct {
	ctx context.Context
	cli *redis.Client
}

/*
NewLock @Editor robotyang at 2023

# NewLock reids锁控制器
*/
func NewLock(ctx context.Context, cli *redis.Client) *lock {
	return &lock{ctx, cli}
}

/*
Lock @Editor robotyang at 2023

# Lock 对key进行上锁
*/
func (r *lock) Lock(key string, expired time.Duration) (bool, error) {
	res, err := r.cli.SetNX(r.ctx, key, 1, expired).Result()
	if err != nil {
		return false, err
	}
	return res, nil
}

/*
Unlock @Editor robotyang at 2023

# Unlock 对key进行解锁（通过 del 命令），key不存在是返回0
*/
func (r *lock) Unlock(key string) (int64, error) {
	res, err := r.cli.Del(r.ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return res, nil
}

/*
Unlock @Editor robotyang at 2023

# Unlock 对key进行解锁（通过 lua 脚本），key不存在则返回0
*/
func (r *lock) UnlockScript(key string) (interface{}, error) {
	lua := `
		if redis.call("get", KEYS[1]) == ARGV[1] then
			return redis.call("del", KEYS[1])
		else
			return 0
		end;
	`
	res, err := r.cli.Eval(r.ctx, lua, []string{key}, 1).Result()
	if err != nil {
		return nil, err
	} else {
		return res, nil
	}
}
