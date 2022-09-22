package main

import (
	"math/rand"
	"os"
	"time"

	// "k8s.io/component-base/cli"
	"k8s.io/component-base/logs"

	"isula.org/qosguardian/cmd/qosguardian/app"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	command := app.NewQoSGuardianCommand()

	logs.InitLogs()
	defer logs.FlushLogs()

	// code := cli.Run(command)
	// os.Exit(code)

	// TODO: init logger

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
