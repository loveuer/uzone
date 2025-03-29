package api_nf

import (
	"bufio"
	"context"
	"io"
	"mime/multipart"

	"github.com/loveuer/nf"
	"github.com/loveuer/uzone/internal/bytesconv"
	"github.com/loveuer/uzone/pkg/cache"
	"github.com/loveuer/uzone/pkg/db"
	"github.com/loveuer/uzone/pkg/es"
	"github.com/loveuer/uzone/pkg/interfaces"
	"github.com/loveuer/uzone/pkg/log"
	"github.com/loveuer/uzone/pkg/mq"
	"github.com/loveuer/uzone/pkg/uapi"
	"gorm.io/gorm"
)

type Ctx struct {
	ctx    *nf.Ctx
	engine uapi.Engine
}

func (c *Ctx) ReqAndResp() (any, any) {
	return c.ctx.Request, c.ctx.Writer
}

func (c *Ctx) Redirect(status int, path string) error {
	return c.ctx.Redirect(path, status)
}

func (c *Ctx) Body() []byte {
	return c.ctx.Body()
}

func (c *Ctx) SetContext(ctx context.Context) {
	c.ctx.SetContext(ctx)
}

func (c *Ctx) Cookies(key string, defaultValue ...string) string {
	return c.ctx.Cookies(key, defaultValue...)
}

func (c *Ctx) FormValue(key string, defaultValue ...string) string {
	return c.ctx.FormValue(key, defaultValue...)
}

func (c *Ctx) JSON(data any) error {
	return c.ctx.JSON(data)
}

func (c *Ctx) Method(override ...string) string {
	return c.ctx.Method(override...)
}

func (c *Ctx) Path(override ...string) string {
	return c.ctx.Path(override...)
}

func (c *Ctx) Scheme() string {
	return c.ctx.Scheme()
}

func (c *Ctx) Protocol() string {
	return c.ctx.Protocol()
}

func (c *Ctx) Query(key string, defaultValue ...string) string {
	return c.ctx.Query(key, defaultValue...)
}

func (c *Ctx) Queries() map[string]string {
	return c.ctx.Queries()
}

func (c *Ctx) SaveFile(fileheader *multipart.FileHeader, path string) error {
	return c.ctx.SaveFile(fileheader, path)
}

func (c *Ctx) SendStatus(status int) error {
	return c.ctx.SendStatus(status)
}

func (c *Ctx) SendString(s string) error {
	return c.ctx.SendString(s)
}

func (c *Ctx) SendStream(stream io.Reader, size ...int) error {
	return c.ctx.SendStream(stream, size...)
}

func (c *Ctx) SendStreamWriter(streamWriter func(*bufio.Writer)) error {
	return c.ctx.SendStreamWriter(streamWriter)
}

func (c *Ctx) Status(status int) uapi.Context {
	c.ctx = c.ctx.Status(status)
	return c
}

func (c *Ctx) Writef(f string, data ...any) (int, error) {
	return c.ctx.Writef(f, data...)
}

func (c *Ctx) WriteString(s string) (int, error) {
	return c.ctx.Write(bytesconv.StringToBytes(s))
}

func (c *Ctx) Drop() error {
	return c.ctx.Drop()
}

func (c *Ctx) UseZone() interfaces.Uzone {
	zone, _ := c.engine.GetUZone()
	return zone
}

func (c *Ctx) UseLogger() *log.UzoneLogger {
	return c.engine.UseLogger()
}

func (c *Ctx) UseDB(opts ...db.SessionOpt) *gorm.DB {
	return c.engine.UseDB()
}

func (c *Ctx) UseCache() cache.Cache {
	return c.engine.UseCache()
}

func (c *Ctx) UseES() *es.Client {
	return c.engine.UseES()
}

func (c *Ctx) UseMQ() *mq.Client {
	return c.engine.UseMQ()
}

func (c *Ctx) BodyParser(out any) error {
	return c.ctx.BodyParser(out)
}

func (c *Ctx) Context() context.Context {
	return c.ctx.Context()
}

func (c *Ctx) FormFile(key string) (*multipart.FileHeader, error) {
	return c.ctx.FormFile(key)
}

func (c *Ctx) GetHeader(key string) string {
	return c.ctx.Get(key)
}

func (c *Ctx) SetHeader(key string, val string) {
	c.ctx.SetHeader(key, val)
}

func (c *Ctx) IP() string {
	return c.ctx.IP()
}

func (c *Ctx) Json(data any) error {
	return c.ctx.JSON(data)
}

func (c *Ctx) Locals(key string, data ...any) any {
	return c.ctx.Locals(key, data...)
}

func (c *Ctx) MultipartForm() (*multipart.Form, error) {
	return c.ctx.MultipartForm()
}

func (c *Ctx) Param(key string) string {
	return c.ctx.Param(key)
}

func (c *Ctx) Next() error {
	return c.ctx.Next()
}

func (c *Ctx) QueryParser(out any) error {
	return c.ctx.QueryParser(out)
}

func (c *Ctx) Write(bytes []byte) (int, error) {
	return c.ctx.Write(bytes)
}

func NewCtx(c *nf.Ctx, engine uapi.Engine) uapi.Context {
	return &Ctx{ctx: c, engine: engine}
}
