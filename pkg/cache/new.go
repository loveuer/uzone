package cache

import (
	"fmt"
	"gitea.com/loveuer/gredis"
	"github.com/go-redis/redis/v8"
	"github.com/loveuer/uzone/pkg/tool"
	"net/url"
)

func _new(uri string) (Cache, error) {
	var (
		client Cache
		err    error
		ins    *url.URL
	)

	if ins, err = url.Parse(uri); err != nil {
		return nil, err
	}

	switch ins.Scheme {
	case "memory", "mem":
		gc := gredis.NewGredis(1024 * 1024)
		client = &_mem{client: gc}
	case "lru":
		if client, err = newLRUCache(); err != nil {
			return nil, err
		}
	case "redis":
		var (
			err error
		)

		addr := ins.Host
		username := ins.User.Username()
		password, _ := ins.User.Password()

		var rc *redis.Client
		rc = redis.NewClient(&redis.Options{
			Addr:     addr,
			Username: username,
			Password: password,
		})

		if err = rc.Ping(tool.Timeout(5)).Err(); err != nil {
			return nil, fmt.Errorf("cache.Init: redis ping err: %s", err.Error())
		}

		client = &_redis{client: rc}
	default:
		return nil, fmt.Errorf("cache type %s not support", ins.Scheme)
	}

	return client, nil
}
