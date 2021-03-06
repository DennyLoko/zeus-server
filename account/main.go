package account

import (
	"github.com/Sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func Run(args map[string]interface{}) {
	logrus.SetLevel(logrus.DebugLevel)

	l := NewServer()

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT)

	err := l.Run()

	if err != nil {
		return
	}

	for s := range sig {
		if s == syscall.SIGINT {
			l.Close()
			os.Exit(0)
		}
	}
}
