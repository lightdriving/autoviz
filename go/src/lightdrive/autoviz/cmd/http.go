package cmd

import (
	"github.com/gobestsdk/gobase/log"
	"lightdrive/autoviz/router"
	"os"
	"os/signal"
	"syscall"
)

func Http() {
	log.Setlogfile("fs4d.log")

	defer func() {
		if error := recover(); error != nil {
			log.Fatal(log.Fields{"panic": error})
			exit(-1)
		}
	}()

	go func() {
		singals := make(chan os.Signal)
		signal.Notify(singals,
			syscall.SIGTERM,
			syscall.SIGINT,
			syscall.SIGKILL,
			syscall.SIGHUP,
			syscall.SIGQUIT,
		)
		<-singals
		exit(0)
	}()

	router.Init()
	router.Run()
}

func exit(status int) {
	log.Info(log.Fields{"app": status})
	router.Stop()
	os.Exit(status)
}
