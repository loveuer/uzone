package uzone

import (
	"os"
	"testing"
)

func Test_bindEnv(t *testing.T) {
	os.Setenv("UZONE.LISTEN.HTTP", "0.0.0.0:99")
	os.Setenv("UZONE.DEBUG", "true")
	p := &_property{}
	bindEnv(p)

	t.Logf("listen.http = %s", p.Listen.Http)
	t.Logf("debug = %v", p.Debug)
}
