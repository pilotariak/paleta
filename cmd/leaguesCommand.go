// Copyright (C) 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/pilotariak/paleta/leagues"
	_ "github.com/pilotariak/paleta/leagues/ctpb"
	_ "github.com/pilotariak/paleta/leagues/ffpb"
	_ "github.com/pilotariak/paleta/leagues/lbpb"
	_ "github.com/pilotariak/paleta/leagues/lcapb"
	_ "github.com/pilotariak/paleta/leagues/lidfpb"
)

// LeaguesCommand is the command which display available leagues
var LeaguesCommand = cli.Command{
	Name: "leagues",
	Subcommands: []cli.Command{
		leaguesListCommand,
	},
}

var leaguesListCommand = cli.Command{
	Name:  "list",
	Usage: "List all leagues",
	Action: func(context *cli.Context) error {

		fmt.Println("Leagues:")
		for _, name := range leagues.ListLeagues() {
			fmt.Printf("- %s\n", name)
		}
		return nil
	},
}
