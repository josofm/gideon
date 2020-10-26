package main

import (
	"flag"
	"fmt"

	"github.com/josofm/gideon/controller"
	"github.com/josofm/gideon/repository"

	"github.com/josofm/gideon/api"
	"github.com/subosito/gotenv"
)

var (
	// version is set at build time
	Version = "No version provided at build time"
)

func init() {
	gotenv.Load()
}

func main() {

	version := false
	flag.BoolVar(&version, "version", false, "Show version")

	flag.Parse()
	if version {
		fmt.Printf("version: %s\n", Version)
		return
	}

	r, err := repository.NewRepository()
	if err != nil {
		panic(err)
	}

	c := controller.NewController(r)

	_ = api.NewApi(c).StartServer()

}
