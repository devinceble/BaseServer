package main

import (
	"log"
	"net/http"
	"os"

	"github.com/devinceble/BaseServer/Helpers"
	"github.com/devinceble/BaseServer/Migrations"
	"github.com/devinceble/BaseServer/Routes"
	"github.com/devinceble/BaseServer/Sockets"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/igm/sockjs-go.v2/sockjs"
)

func main() {
	var port string
	port = ":8000"
	if len(os.Args) >= 2 {
		if len(os.Args) >= 3 {
			port = ":" + os.Args[2]
		}
		switch os.Args[1] {
		case "start":
			r := Routes.Route()
			http.Handle("/", r)
			http.Handle("/echo/", sockjs.NewHandler("/echo", sockjs.DefaultOptions, Sockets.Chat))
			Helpers.BaseInfo("PORT"+port, "BASESERVER STARTED")
			log.Fatal(http.ListenAndServe(port, nil))
			break
		case "migrate":
			Migrations.Migrate()
			break
		default:
			cmd()
		}

	} else {
		cmd()
	}
}

func cmd() {
	Helpers.BaseInfo("Usage:", "BaseServer <command>\n ")
	Helpers.BaseInfo("BASE SERVER STRUCTURE", "\n ")
	Helpers.BaseInfo("HELP", "#To Start Server --> BaseServer server PORT\n ")
	Helpers.BaseInfo("HELP", "$ BaseServer start 8000\n ")
	Helpers.BaseInfo("HELP", "#To Migrate DATABASE --> BaseServer migrate\n ")
	Helpers.BaseInfo("HELP", "$ BaseServer migrate\n\n ")
}
