package main

import (
	"github.com/kubernetix/k8x/v1/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"runtime"
)

func init() {
	runtime.LockOSThread()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	cmd.Execute()
}
