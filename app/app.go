package app

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"fmt"
	"runtime/debug"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/exceptions"
	"github.com/thuongtruong1009/zoomer/internal/adapter"
)

func Run() {
	runtime.GC()
	debug.FreeOSMemory()

	numProcs := runtime.NumCPU()
	runtime.GOMAXPROCS(numProcs)

	adt := adapter.Adapter()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(interrupt, os.Kill)

	select {
	case s := <-interrupt:
		exceptions.SystemLog(fmt.Sprintf("%s: %s", constants.ServerGotTerminate, s.String()))
	case err := <-adt.Notify():
		exceptions.SystemLog(fmt.Sprintf("%s: %s", constants.ServerGotError, err.Error()))
	}

	if err := adt.Shutdown(); err != nil {
		exceptions.Fatal(constants.ErrorShuttdownServer, fmt.Sprintf("%v", err))
	}
}
