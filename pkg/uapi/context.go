package uapi

import (
	"bufio"
	"context"
	"io"
	"mime/multipart"

	"github.com/loveuer/uzone/pkg/cache"
	"github.com/loveuer/uzone/pkg/db"
	"github.com/loveuer/uzone/pkg/es"
	"github.com/loveuer/uzone/pkg/interfaces"
	"github.com/loveuer/uzone/pkg/log"
	"github.com/loveuer/uzone/pkg/mq"
	"gorm.io/gorm"
)

type Context interface {

	// ReqAndResp
	//   - for api engine based on http, will return *http.Request and http.ResponseWriter
	//   - for api engine based on fasthttp, will return *fasthttp.Request, *fasthttp.Response
	ReqAndResp() (any, any)

	// App() *App

	Body() []byte
	SetContext(ctx context.Context)
	Cookies(key string, defaultValue ...string) string
	//Request() *fasthttp.Request
	//Response() *fasthttp.Response
	FormValue(key string, defaultValue ...string) string
	JSON(data any) error
	//XML(data any) error
	Method(override ...string) string
	//ClientHelloInfo() *tls.ClientHelloInfo
	Path(override ...string) string
	Scheme() string
	Protocol() string
	Query(key string, defaultValue ...string) string
	Queries() map[string]string
	SaveFile(fileheader *multipart.FileHeader, path string) error
	SendStatus(status int) error
	SendString(body string) error
	SendStream(stream io.Reader, size ...int) error
	SendStreamWriter(streamWriter func(*bufio.Writer)) error
	Status(status int) Context
	Writef(f string, a ...any) (int, error)
	WriteString(s string) (int, error)
	Drop() error
	Redirect(status int, location string) error

	UseZone() interfaces.Uzone
	UseLogger() *log.UzoneLogger
	UseDB(opts ...db.SessionOpt) *gorm.DB
	UseCache() cache.Cache
	UseES() *es.Client
	UseMQ() *mq.Client
	BodyParser(out any) error
	Context() context.Context
	FormFile(string) (*multipart.FileHeader, error)
	GetHeader(string) string
	SetHeader(string, string)
	IP() string
	Json(any) error
	Locals(key string, data ...any) any

	// get method or rewrite method
	MultipartForm() (*multipart.Form, error)
	Param(string) string

	// get path or rewrite path
	Next() error
	QueryParser(any) error

	Write([]byte) (int, error)
}
