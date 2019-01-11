// Copyright (C) 2016-2019 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli"

	"github.com/pilotariak/paleta/pkg/cmd"
	_ "github.com/pilotariak/paleta/pkg/leagues/ctpb"
	_ "github.com/pilotariak/paleta/pkg/leagues/ffpb"
	_ "github.com/pilotariak/paleta/pkg/leagues/lbpb"
	_ "github.com/pilotariak/paleta/pkg/leagues/lcapb"
	_ "github.com/pilotariak/paleta/pkg/leagues/lidfpb"
	_ "github.com/pilotariak/paleta/pkg/logging"
	"github.com/pilotariak/paleta/pkg/version"
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
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		} else {
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err)
	}
}
