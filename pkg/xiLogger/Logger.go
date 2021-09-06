package xiLogger


import (
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var Log logrus.Logger

func init()  {
	Log  = logrus.Logger{}

	Log.SetFormatter(&logrus.TextFormatter{DisableColors: false})
	Log.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	Log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	Log.SetLevel(logrus.DebugLevel)
}

func Logger() *logrus.Entry {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("Could not get context info for logger!")
	}

	filename := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	funcname := runtime.FuncForPC(pc).Name()
	fn := funcname[strings.LastIndex(funcname, ".")+1:]
	return logrus.WithField("file", filename).WithField("function", fn)
}

func IsDebugEnabled() bool  {
	if Log.IsLevelEnabled(logrus.DebugLevel){
		return true
	}
	return false
}