package main

import (
	"flag"

	"github.com/fengde/gocommon/confx"
	"github.com/fengde/gocommon/confx/example/config"
	"github.com/fengde/gocommon/logx"
	_ "github.com/joho/godotenv/autoload"
)

var configFile = flag.String("f", "etc/example-api.yaml", "the config file")

func main() {
	flag.Parse()
	var c config.Config
	confx.MustLoad(*configFile, &c, confx.UseEnv())

	logx.Info(c.Name)
}
