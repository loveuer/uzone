package uzone

import (
	"fmt"
	"github.com/loveuer/uzone/pkg/log"
	"github.com/spf13/cast"
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
	"time"

	"github.com/loveuer/uzone/pkg/tool"
)

type _property struct {
	Debug  bool `yaml:"debug" json:"debug" env:"UZONE.DEBUG"`
	Listen struct {
		Http string `yaml:"http" json:"http" env:"UZONE.LISTEN.HTTP"`
	} `yaml:"listen" json:"listen"`
	DB struct {
		URI string `json:"uri" env:"UZONE.DB.URI"`
	} `yaml:"db" json:"db"`
	Cache struct {
		URI string `json:"uri" env:"UZONE.CACHE.URI"`
	} `yaml:"cache" json:"cache"`
	Elasticsearch struct {
		URI string `yaml:"uri" json:"uri" env:"UZONE.ELASTICSEARCH.URI"`
	} `yaml:"elasticsearch" json:"elasticsearch"`
}

var property = &_property{}

func init() {
	time.Local = time.FixedZone("CST", 8*3600)

	var (
		err error
		bs  []byte

		configFn = func(path string) error {
			if bs, err = os.ReadFile(path); err != nil {
				log.New().Debug("[%30s] read %s err, err = %s", "init", path, err.Error())
				return err
			}

			if err = yaml.Unmarshal(bs, property); err != nil {
				log.New().Debug("[%30s] unmarshal %s err, err = %s", "init", path, err.Error())
				return err
			}

			return nil
		}
	)

	if err = configFn("etc/config.yaml"); err == nil {
		goto BindEnv
	}

	if err = configFn("etc/config.yml"); err == nil {
		goto BindEnv
	}

	_ = configFn("etc/config.json")

BindEnv:
	_ = bindEnv(property)

	if property.Debug {
		tool.TablePrinter(property)
	}
}

func bindEnv(data any) error {
	rv := reflect.ValueOf(data)

	if rv.Type().Kind() != reflect.Pointer {
		return fmt.Errorf("can only bind ptr")
	}

	rv = rv.Elem()

	return bindStruct(rv)
}

func bindStruct(rv reflect.Value) error {
	if rv.Type().Kind() != reflect.Struct {
		return fmt.Errorf("can only bind struct ptr")
	}

	for i := 0; i < rv.NumField(); i++ {
		f := rv.Field(i)

		if f.Type().Kind() == reflect.Pointer {
			f = f.Elem()
		}

		if f.Type().Kind() == reflect.Struct {
			return bindStruct(f)
		}

		if !f.CanSet() {
			continue
		}

		tag := rv.Type().Field(i).Tag.Get("env")
		if tag == "" || tag == "-" {
			continue
		}

		bv := os.Getenv(tag)
		if bv == "" {
			continue
		}

		switch f.Type().Kind() {
		case reflect.String:
			f.SetString(bv)
		case reflect.Bool:
			f.SetBool(cast.ToBool(bv))
		case reflect.Int64, reflect.Int, reflect.Uint64, reflect.Uint, reflect.Int32, reflect.Uint32, reflect.Int16, reflect.Uint16, reflect.Int8, reflect.Uint8:
			f.SetInt(cast.ToInt64(bv))
		case reflect.Float64, reflect.Float32:
			f.SetFloat(cast.ToFloat64(bv))
		default:
			log.New().Warn("[%30s] unsupported env binding, type = %s, value = %s", "init", f.Type().Kind().String(), bv)
		}
	}

	return nil
}
