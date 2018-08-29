// Code generated by TKP Cli. DO NOT EDIT.
package main

import (
	"flag"
	"os"

	"github.com/tokopedia/tdk/go/app"
	"github.com/tokopedia/tdk/go/flags"
	"github.com/tokopedia/tdk/go/log"
	"punyasaya/pkg/server"
)

type serviceFlags struct {
	TkpConfigPath  string
	SidecarAddress string
	Test           bool
}

func (sf *serviceFlags) Parse(fs *flag.FlagSet, args []string) error {
	fs.StringVar(&sf.TkpConfigPath, "config_path", "", "tkp application config path")
	fs.StringVar(&sf.SidecarAddress, "sidecar_address", "info", "sidecar address")
	fs.BoolVar(&sf.Test, "t", false, "testing flag, to test application")
	return fs.Parse(args)
}

var sf serviceFlags

func main() {
	flags.Parse(&sf)

	punyasaya, err := app.New(app.Options{
		Name:       "punyasaya",
		TeamName:   "devcamp",
		ConfigPath: sf.TkpConfigPath,
	})

	if err != nil {
		log.Fatal(err)
	}

	err = server.Init(punyasaya)
	if err != nil {
		log.Fatal(err)
	}

	// Please don't remove pattern below
	//@@SERVICE_INIT@@//
	httpServer := server.NewHttpServer()
	if err := punyasaya.RegisterHTTPServer(httpServer); err != nil {
		log.Fatal(err)
	}

	if sf.Test {
		log.Println("Test complete, exiting program")
		os.Exit(0)
	}

	if err := punyasaya.Run(); err != nil {
		log.Fatal(err)
	}
}
