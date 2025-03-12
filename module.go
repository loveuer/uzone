package uzone

import (
	"github.com/loveuer/uzone/pkg/api"
	"github.com/loveuer/uzone/pkg/log"
	"github.com/loveuer/uzone/pkg/mq"

	"github.com/loveuer/uzone/pkg/cache"
	"github.com/loveuer/uzone/pkg/db"
	"github.com/loveuer/uzone/pkg/es"
	"github.com/loveuer/uzone/pkg/interfaces"
)

type module func(u *uzone)

func InitDB(opts ...db.OptionFn) module {
	db, err := db.New(opts...)
	if err != nil {
		log.New().Panic(err.Error())
	}

	return func(u *uzone) {
		u.db = db
	}
}

func InitCache(opts ...cache.OptionFn) module {
	cache, err := cache.New(opts...)
	if err != nil {
		log.New().Panic(err.Error())
	}

	return func(u *uzone) {
		u.cache = cache
	}
}

func InitES(opts ...es.OptionFn) module {
	client, err := es.New(opts...)
	if err != nil {
		log.New().Panic(err.Error())
	}

	return func(u *uzone) {
		u.es = client
	}
}

func InitMQ(opts ...mq.OptionFn) module {
	client, err := mq.New(opts...)
	if err != nil {
		log.New().Panic(err.Error())
	}

	return func(u *uzone) {
		u.mq = client
	}
}

func InitApi(engine api.Engine, cfgs ...api.Config) module {
	cfg := api.Config{}
	if len(cfgs) > 0 {
		cfg = cfgs[0]
	}

	if cfg.Address == "" {
		cfg.Address = property.Listen.Http
	}

	if cfg.Address == "" {
		cfg.Address = "localhost:8080"
	}

	return func(u *uzone) {
		engine.SetUZone(u, cfg)
		u.api = engine
	}
}

func InitTaskChan(ch <-chan func(uzone interfaces.Uzone) error) module {
	return func(u *uzone) {
		if u.taskCh == nil {
			u.taskCh = make([]<-chan func(u interfaces.Uzone) error, 0)
		}

		u.taskCh = append(u.taskCh, ch)
	}
}

// sync functions
// 添加 同步执行函数
func InitFn(fns ...func(interfaces.Uzone)) module {
	return func(u *uzone) {
		u.initFns._sync = append(u.initFns._sync, fns...)
	}
}

// async functions
// 添加 异步执行函数
func InitAsyncFn(fns ...func(interfaces.Uzone)) module {
	return func(u *uzone) {
		u.initFns._async = append(u.initFns._async, fns...)
	}
}
