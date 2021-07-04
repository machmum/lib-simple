package liblog

import "github.com/sirupsen/logrus"

var skip bool

// Logger indicates minimal method to implement logger
type Logger interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Fatal(...interface{})
}

// logger will be an extendable method from logrus.Entry
type logger struct {
	f   logrus.Fields
	opt Options
}

func (l *logger) Debug(args ...interface{}) {
	if !skip {
		e := logrus.WithFields(l.f)
		e.Debug(args...)
	}
}

func (l *logger) Info(args ...interface{}) {
	if !skip {
		e := logrus.WithFields(l.f)
		e.Info(args...)
	}
}

func (l *logger) Warn(args ...interface{}) {
	if !skip {
		e := logrus.WithFields(l.f)
		e.Warn(args...)
	}
}

func (l *logger) Error(args ...interface{}) {
	if !skip {
		e := logrus.WithFields(l.f)
		e.Error(args...)
	}
}

func (l *logger) Fatal(args ...interface{}) {
	if !skip {
		e := logrus.WithFields(l.f)
		e.Fatal(args...)
	}
}
