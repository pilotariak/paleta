// Copyright (C) 2016, 2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	// "fmt"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"

	_ "github.com/pilotariak/paleta/leagues/ctpb"
	_ "github.com/pilotariak/paleta/leagues/ffpb"
	_ "github.com/pilotariak/paleta/leagues/lbpb"
	_ "github.com/pilotariak/paleta/leagues/lcapb"
	_ "github.com/pilotariak/paleta/leagues/lidfpb"

	"github.com/pilotariak/paleta/cmd"
	"github.com/pilotariak/paleta/version"
)

func main() {
	app := cli.NewApp()
	app.Name = "paleta"
	app.Usage = "CLI for Pelota news"
	app.Version = version.Version

	app.Commands = []cli.Command{
		cmd.VersionCommand,
		cmd.LeaguesCommand,
		cmd.LeagueCommand,
	}

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "debug",
			// Value: false,
			Usage: "Enable debug mode",
		},
	}
	app.Action = func(context *cli.Context) error {
		if context.Bool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		} else {
			logrus.SetLevel(logrus.WarnLevel)
		}
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
