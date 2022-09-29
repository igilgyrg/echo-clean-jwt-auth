package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
)

type logger struct {
	*logrus.Entry
}

type Log interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

func newLogger() Log {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s %s", filename, frame.Line), fmt.Sprintf("%s", frame.Function)
		},
		DisableColors: false,
		FullTimestamp: true,
	}

	l.SetOutput(os.Stdout)

	for _, level := range logrus.AllLevels {
		l.SetLevel(level)
	}

	return &logger{logrus.NewEntry(l)}
}
