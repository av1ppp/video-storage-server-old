package main

import (
	"fmt"
	"os"

	"github.com/AviParampampam/video-storage-server/config"
	"github.com/AviParampampam/video-storage-server/server"
	"github.com/fatih/color"
)

func main() {
	conf := &config.Config{}
	err := config.ParseFile("config.yaml", conf)
	if err != nil {
		exitError(err)
	}
	err = config.ParseEnv(conf)
	if err != nil {
		exitError(err)
	}

	srv := server.New(conf)
	if err := srv.Listen(); err != nil {
		exitError(err)
	}
}

func exitError(err error) {
	fmt.Println(color.RedString(err.Error()))
	os.Exit(1)
}
