package app

import (
	"fmt"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs"
	"github.com/thuongtruong1009/zoomer/infrastructure/configs/parameter"
	"github.com/thuongtruong1009/zoomer/internal/server/adapter"
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

	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGHUP, syscall.SIGQUIT,)
	signal.Notify(c, os.Kill)

	cfg := configs.LoadConfigs(constants.EnvConfPath)
	paramCfg := parameter.LoadParameterConfigs(constants.ParamConfPath)

	adt := adapter.NewAdapter(cfg, paramCfg)

	select {
	case s := <-c:
		exceptions.SystemLog(fmt.Sprintf("%s: %s", constants.ServerGotTerminate, s.String()))
	case err := <-adt.Notify():
		exceptions.SystemLog(fmt.Sprintf("%s: %s", constants.ServerGotError, err.Error()))
	}

	exceptions.SystemLog(constants.ServerShutdown)
	ctx, cancel := context.WithTimeout(context.Background(), paramCfg.ServerConf.ShutdownTimeout * time.Second)
	exceptions.SystemLog(constants.ServerExitedProperly)

	p, _ := os.FindProcess(os.Getpid())
	p.Signal(syscall.SIGINT)

	defer cancel()

	if err := adt.Shutdown(ctx); err != nil {
		exceptions.Fatal(constants.ErrorShuttdownServer, fmt.Sprintf("%v", err))
	}
}
