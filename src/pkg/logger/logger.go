package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func Logger() *logrus.Logger {

	f, err := os.OpenFile(fmt.Sprintf("logs/%s.log", time.Now().Format("2006-01-02")), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic("where is log file?")
	}
	l := logrus.New()
	l.SetOutput(f)
	return l
}
