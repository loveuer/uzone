package db

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/glebarez/sqlite"
	"github.com/loveuer/uzone/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(opts ...OptionFn) (*Client, error) {
	var (
		err    error
		client = &Client{
			uri: "sqlite://data.db",
		}
	)

	for _, fn := range opts {
		fn(client)
	}

	if client.tx, err = _new(client.uri); err != nil {
		return nil, err
	}

	if len(client.models) > 0 {
		if err = client.Session().AutoMigrate(client.models...); err != nil {
			return nil, err
		}
	}

	return client, nil
}

func _new(uri string) (*gorm.DB, error) {
	ins, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	var (
		username = ""
		password = ""
		tx       *gorm.DB
	)

	if ins.User != nil {
		username = ins.User.Username()
		password, _ = ins.User.Password()
	}

	switch ins.Scheme {
	case "sqlite":
		path := strings.TrimPrefix(uri, ins.Scheme+"://")
		log.New().Debug("db.New: type = %s, path = %s", ins.Scheme, path)
		tx, err = gorm.Open(sqlite.Open(path))
	case "mysql", "mariadb":
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", username, password, ins.Host, ins.Path, ins.RawQuery)
		log.New().Debug("db.New: type = %s, dsn = %s", ins.Scheme, dsn)
		tx, err = gorm.Open(mysql.Open(dsn))
	case "pg", "postgres", "postgresql":
		opts := make([]string, 0)
		for key, val := range ins.Query() {
			opts = append(opts, fmt.Sprintf("%s=%s", key, val))
		}
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s %s", ins.Hostname(), username, password, ins.Path, ins.Port(), strings.Join(opts, " "))
		log.New().Debug("db.New: type = %s, dsn = %s", ins.Scheme, dsn)
		tx, err = gorm.Open(postgres.Open(dsn))
	default:
		return nil, fmt.Errorf("invalid database type(uri_scheme): %s", ins.Scheme)
	}

	if err != nil {
		return nil, err
	}

	return tx, nil
}
