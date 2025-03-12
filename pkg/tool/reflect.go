package tool

import (
	"reflect"
	"runtime"
)

func GetFunctionName(i interface{}) string {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Func {
		return "<not a function>"
	}
	pc := v.Pointer()
	if pc == 0 {
		return "<nil function>"
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "<unknown function>"
	}
	return fn.Name()
}
