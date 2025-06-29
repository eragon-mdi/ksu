package server

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitingForShutdownSignal() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)

	<-done
}
