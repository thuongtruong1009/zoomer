package app

import (
	"fmt"
	"github.com/thuongtruong1009/zoomer/configs"
	"github.com/thuongtruong1009/zoomer/configs/parameter"
	"github.com/thuongtruong1009/zoomer/internal/adapter"
	"github.com/thuongtruong1009/zoomer/pkg/constants"
	"github.com/thuongtruong1009/zoomer/pkg/exceptions"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"syscall"
	"time"
	"context"
)

func Run() {
	runtime.GC()
	debug.FreeOSMemory()

	numProcs := runtime.NumCPU()
	runtime.GOMAXPROCS(numProcs)

	cfg := configs.LoadConfigs(constants.EnvConfPath)
	paramCfg := parameter.LoadParameterConfigs(constants.ParamConfPath)

	adt := adapter.Adapter(cfg, paramCfg)

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

	exceptions.SystemLog(constants.ServerShutdown)
	ctx, cancel := context.WithTimeout(context.Background(), paramCfg.ServerConf.ShutdownTimeout * time.Second)
	exceptions.SystemLog(constants.ServerExitedProperly)

	defer cancel()

	if err := adt.Shutdown(ctx); err != nil {
		exceptions.Fatal(constants.ErrorShuttdownServer, fmt.Sprintf("%v", err))
	}
}
