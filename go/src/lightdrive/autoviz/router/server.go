package router

import (
	"github.com/gobestsdk/gobase/httpserver"
	"lightdrive/autoviz/status"
)

var (
	App = httpserver.New(status.Appname)
)

func Run() {

	App.SetPort(status.Port)
	App.Run()
}

func Stop() {

	App.Stop()
}
