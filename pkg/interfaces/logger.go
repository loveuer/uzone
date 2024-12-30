package interfaces

type Logger interface {
	With(args ...interface{}) Logger
	Debug(string, ...any)
	Info(string, ...any)
	Warn(string, ...any)
	Error(string, ...any)
	Panic(string, ...any)
	Fatal(string, ...any)
}
