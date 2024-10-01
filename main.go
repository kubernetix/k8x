package main

import (
	"github.com/kubernetix/k8x/v1/cmd"
	"github.com/kubernetix/k8x/v1/internal/dotenv"
	"runtime"
)

func init() {
	runtime.LockOSThread()
	dotenv.Load()
}

func main() {
	cmd.Execute()
}
