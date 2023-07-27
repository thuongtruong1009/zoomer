package main

import (
	"github.com/thuongtruong1009/zoomer/configs"
	"github.com/thuongtruong1009/zoomer/internal/app"
)

func main() {
	cfg := configs.NewConfig()

	app.Run(cfg)
}
