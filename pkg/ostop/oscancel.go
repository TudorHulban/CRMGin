package ostop

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// NewOSCancellableCtx Returns an OS cancellable context.
func NewOSCancellableCtx(signals ...os.Signal) (context.Context, context.CancelFunc) {
	if len(signals) == 0 {
		signals = append(signals, syscall.SIGTERM)
		signals = append(signals, syscall.SIGINT)
	}

	ctx, cancel := context.WithCancel(context.Background())

	chQuit := make(chan os.Signal, 1)
	signal.Notify(chQuit, signals...)

	go func() {
		<-chQuit
		cancel()
	}()

	return ctx, cancel
}
