# UZone

# make golang BE easy

### install in your project

- `go get github.com/loveuer/uzone@latest`
  
### Usage

- 1. [examples](https://github.com/loveuer/uzone/tree/master/example)

- 2. usage present
  ```go
  app := uzone.New()

  app.With(db, es, api)

  app.Run(ctx)
  ```

- 3. simple example
    ```go
  package main

  import (
      "github.com/loveuer/uzone"
      "github.com/loveuer/uzone/pkg/api"
      "github.com/loveuer/uzone/pkg/db"
      "github.com/loveuer/uzone/pkg/interfaces"
      "net/http"
      "time"
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
      app.With(uzone.InitES())
      app.With(uzone.InitMQ())
      app.With(uzone.InitApi(api.New()))

      app.With(uzone.InitFn(func(u interfaces.Uzone) {
          u.UseLogger().Debug("[init] create init record")
          u.UseDB().Create(&Record{Key: "init"})
      }))

      app.With(uzone.InitAsyncFn(func(u interfaces.Uzone) {
          time.Sleep(10 * time.Second)
          u.UseLogger().Info("async: run!!!")
      }))

      var ch = make(chan func(u interfaces.Uzone) error)
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

      app.GET("/api/available", func(c *api.Ctx) error {
          c.UseLogger().With("module", "example").Warn("hello world")
          time.Sleep(500 * time.Millisecond)
          return c.Status(500).SendString("hello world")
      })

      app.POST("/api/record/create", func(c *api.Ctx) error {
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
    ```

### Config
  > The program will load configuration files in the following order:
  - `etc/config.yaml`
  - `etc/config.yml`
  - `config.json`
  > Environment variables will take precedence and override any matching configurations found in the files above.

  > envs:
  - `UZONE.DEBUG=true`
  - `UZONE.LISTEN.HTTP=0.0.0.0:80`

  > etc/config.yaml:
  - ```yaml
    ```