package main

import (
	"github.com/kubernetix/k8x/v1/cmd"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	cmd.Execute()
}
