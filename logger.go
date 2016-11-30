package lowger

func init() {
	// add default plugin
	AddPlugin(&ConsoleLogger{})
}

func Error(messages ...interface{}) {
	go runErrorPlugins(messages...)
}

func Warning(messages ...interface{}) {
	go runWarningPlugins(messages...)
}

func Info(messages ...interface{}) {
	go runInfoPlugins(messages...)
}

func Fatal(messages ...interface{}) {
	go runFatalPlugins(messages...)
}

type Logger interface {
	Error(messages ...interface{})
	Warning(messages ...interface{})
	Info(messages ...interface{})
	Fatal(messages ...interface{})
}

var Plugins []Logger

func AddPlugin(l Logger) int {
	Plugins = append(Plugins, l)
	return len(Plugins) - 1
}

func RemovePlugin(idx int) {
	Plugins = append(Plugins[:idx], Plugins[idx+1:]...)[:len(Plugins)-1]
}

func runErrorPlugins(messages ...interface{}) {
	for _, l := range Plugins {
		l.Error(messages...)
	}
}

func runWarningPlugins(messages ...interface{}) {
	for _, l := range Plugins {
		l.Warning(messages...)
	}
}
func runInfoPlugins(messages ...interface{}) {
	for _, l := range Plugins {
		l.Info(messages...)
	}
}

func runFatalPlugins(messages ...interface{}) {
	for _, l := range Plugins {
		l.Fatal(messages...)
	}
}
