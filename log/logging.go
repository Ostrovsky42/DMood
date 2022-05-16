package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
)

func Init()  {
	l:=logrus.New()
	l.SetReportCaller(true)
	l.Formatter=&logrus.TextFormatter{CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
		fileName:=path.Base(frame.File)
		return fmt.Sprintf("%s", frame.Func),fmt.Sprintf("%s,%d", fileName,frame.Line)
	},
		FullTimestamp: true}
}