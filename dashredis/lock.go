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

# NewLock  need redis v6
*/
func NewLock(ctx context.Context, cli *redis.Client) *lock {
	return &lock{ctx, cli}
}

func (r *lock) Lock(key string, expired time.Duration) (bool, error) {
	res, err := r.cli.SetNX(r.ctx, key, time.Now().UnixNano(), expired).Result()
	if err != nil {
		return false, err
	}
	return res, nil
}

func (r *lock) Free(key string) (int64, error) {
	res, err := r.cli.Del(r.ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return res, nil
}

func (r *lock) freeByScript() error {
	//lua := `
	//	if redis.call("get", KEYS[1]) == ARGV[1] then
	//		return redis.call("del", KEYS[1])
	//	else
	//		return 0
	//	end;
	//`
	//res, err := r.cli.Eval(r.ctx, lua, []string{r.key}, r.val).Result()
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("res", res)
	//}
	//return err
	return nil
}
