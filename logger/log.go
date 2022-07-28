package logger

import (
	"github.com/lqgl/backlog"
	"log"
	"sync"
)

var (
	Logger         *backlog.BackLog
	onceInitLogger sync.Once
)

func init() {
	onceInitLogger.Do(func() {
		fileWriter := backlog.NewFileWriter("log", 0, 0, 0)
		l := backlog.NewBackLog(backlog.Debug, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile, backlog.WithFileWriter(fileWriter))
		Logger = l
	})
}
