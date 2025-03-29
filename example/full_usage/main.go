package main

import (
	"crypto/tls"
	"time"

	"github.com/loveuer/uzone"
	"github.com/loveuer/uzone/example/full_usage/controller"
	"github.com/loveuer/uzone/pkg/api"
	"github.com/loveuer/uzone/pkg/cache"
	"github.com/loveuer/uzone/pkg/db"
	"github.com/loveuer/uzone/pkg/es"
	"github.com/loveuer/uzone/pkg/interfaces"
	"github.com/loveuer/uzone/pkg/mq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func init() {

}

func main() {
	app := uzone.New()

	app.With(uzone.InitCache(cache.WithURI("redis://localhost:6379")))

	app.With(uzone.InitMQ(mq.WithURI("amqp://localhost:5672")))

	app.With(uzone.InitDB(
		db.WithURI("sqlite://data.db"),
		db.WithAutoMigrate(&controller.Record{}),
	))

	app.With(uzone.InitES(
		es.WithURI("http://localhost:9200"),
		es.WithPing(3),
	))

	app.With(uzone.InitFn(func(u interfaces.Uzone) {
		u.UseLogger().Info("init fn start...")
		u.UseLogger().Info("%v", u.UseMQ().Publish(u.UseCtx(), "uzone", amqp.Publishing{
			Body: []byte("hello world"),
		}))
	}))

	app.With(uzone.InitAsyncFn(func(u interfaces.Uzone) {
		time.Sleep(10 * time.Second)
		err := u.UseMQ().Publish(u.UseCtx(), "uzone", amqp.Publishing{Body: []byte("hello uzone")})
		u.UseLogger().Info("publish mq msg: %v", err)
	}))

	app.With(uzone.InitAsyncFn(func(u interfaces.Uzone) {
		ch, err := u.UseMQ().Consume(u.UseCtx(), "uzone")
		if err != nil {
			u.UseLogger().Error(err.Error())
			return
		}

		for m := range ch {
			u.UseLogger().Info("consume: got msg = %s", string(m.Body))
			m.Ack(false)
		}
	}))

	app.With(uzone.InitApi(
		controller.New(),
		api.SetListenAddress("0.0.0.0:8080"),
		api.SetTLS(&tls.Config{}),
	))

	app.RunSignal()
}
