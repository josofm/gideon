package main

import (
	"flag"
	"fmt"

	"github.com/josofm/gideon/controller"

	"github.com/josofm/gideon/api"
)

var (
	// version is set at build time
	Version = "No version provided at build time"
)

func main() {

	version := false
	flag.BoolVar(&version, "version", false, "Show version")

	flag.Parse()
	if version {
		fmt.Printf("version: %s\n", Version)
		return
	}

	c := controller.Controller{}

	_ = api.NewApi(c).StartServer()

}
