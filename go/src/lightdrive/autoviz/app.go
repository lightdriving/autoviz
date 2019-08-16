package main

import (
	"github.com/gobestsdk/gobase/log"
	"gopkg.in/alecthomas/kingpin.v2"
	"lightdrive/autoviz/cmd"
	"lightdrive/autoviz/status"
	"os"
	"strconv"
)

var (
	app  = kingpin.New(status.Appname, "")
	http = app.Command("http", "run as a httpserver")

	port  = http.Flag("dist", "dist path").String()

)

func main() {

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case http.FullCommand():

		if port != nil && *port != "" {

			status.Port,_= strconv.Atoi(*port)
			log.Info(log.Fields{
				"port": status.Port,
			})
		}

		cmd.Http()
	}
}
