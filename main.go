package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/drone-plugins/drone-cache/plugin"
	"github.com/urfave/cli"
)

var build = "0" // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "cache plugin"
	app.Usage = "cache plugin"
	app.Version = fmt.Sprintf("1.0.%s", build)
	app.Action = artifactoryPlugin
	app.Flags = append(
		artifactoryFlags,
		plugin.PluginFlags()...
	)
	// app.Commands = []cli.Command{
	// 	artifactoryCmd,
	// }

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
