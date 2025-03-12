package main

import (
	"net/http"
	"time"

	"github.com/loveuer/uzone"
	"github.com/loveuer/uzone/pkg/api"

	api_nf "github.com/loveuer/uzone/pkg/api.nf"
	// api_fiber "github.com/loveuer/uzone/pkg/api.fiber"
	"github.com/loveuer/uzone/pkg/db"
	"github.com/loveuer/uzone/pkg/interfaces"
)

type Record struct {
	Id    uint64 `json:"id" gorm:"primaryKey;column:id"`
	Key   string `json:"key" gorm:"column:key;unique"`
	Value string `json:"value" gorm:"column:value"`
}

func main() {
	app := uzone.New(uzone.Config{Debug: true})

	app.With(uzone.InitCache())
	app.With(uzone.InitDB(db.WithAutoMigrate(&Record{})))
	//app.With(uzone.InitES())
	//app.With(uzone.InitMQ())
	// app.With(uzone.InitApi(api_fiber.New()))
	app.With(uzone.InitApi(api_nf.New()))

	app.With(uzone.InitFn(func(u interfaces.Uzone) {
		u.UseLogger().Debug("[init] create init record")
		u.UseDB().Create(&Record{Key: "init"})
	}))

	app.With(uzone.InitAsyncFn(func(u interfaces.Uzone) {
		time.Sleep(10 * time.Second)
		u.UseLogger().Info("async: run!!!")
	}))

	ch := make(chan func(u interfaces.Uzone) error)
	app.With(uzone.InitTaskChan(ch))
	go func() {
		for {
			time.Sleep(10 * time.Second)
			ch <- func(u interfaces.Uzone) error {
				u.UseLogger().Info("chan worker: run")
				return nil
			}
		}
	}()

	app.GET("/api/available", func(c api.Context) error {
		c.UseLogger().With("module", "example").Warn("hello world")
		time.Sleep(500 * time.Millisecond)
		return c.Status(500).SendString("hello world")
	})

	app.POST("/api/record/create", func(c api.Context) error {
		var (
			err error
			req = new(Record)
		)

		if err = c.BodyParser(req); err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		if err = c.UseDB().Create(req).Error; err != nil {
			return c.Status(http.StatusInternalServerError).SendString(err.Error())
		}

		return c.JSON(req)
	})

	app.RunSignal()
}
