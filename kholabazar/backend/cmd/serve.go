package cmd

import (
	"kholabazar/config"
	"kholabazar/rest"
)

func Serve() {
	conf := config.GetConfig()
	rest.Start(conf)
}
