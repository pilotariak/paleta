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

package cmd

import (
	"fmt"

	"github.com/urfave/cli"

	"github.com/pilotariak/paleta/pkg/leagues"
)

// LeagueCommand is the command which manage a league
var LeagueCommand = cli.Command{
	Name: "league",
	Subcommands: []cli.Command{
		leagueDescribeCommand,
		leagueLevelsCommand,
		leagueDisciplinesCommand,
		leagueChallengesCommand,
		leagueResultsCommand,
	},
}

var leagueDescribeCommand = cli.Command{
	Name:  "describe",
	Usage: "Describe current league",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "league",
			Usage: "the id of the league",
		},
	},
	Action: func(context *cli.Context) error {
		if !context.IsSet("league") {
			return fmt.Errorf("Please specify the id of the league to used via the --league option")
		}
		league, err := leagues.New(context.String("league"))
		if err != nil {
			return err
		}
		leagues.Describe(league)
		return nil
	},
}

var leagueLevelsCommand = cli.Command{
	Name:  "levels",
	Usage: "List all levels",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "league",
			Usage: "the id of the league",
		},
	},
	Action: func(context *cli.Context) error {
		if !context.IsSet("league") {
			return fmt.Errorf("Please specify the id of the league to used via the --league option")
		}
		league, err := leagues.New(context.String("league"))
		if err != nil {
			return err
		}
		fmt.Println("Levels:")
		for k, v := range league.Levels() {
			fmt.Printf("- [%s] %s\n", k, v)
		}
		return nil
	},
}

var leagueDisciplinesCommand = cli.Command{
	Name:  "disciplines",
	Usage: "List all disciplines",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "league",
			Usage: "the id of the league",
		},
	},
	Action: func(context *cli.Context) error {
		if !context.IsSet("league") {
			return fmt.Errorf("Please specify the id of the league to used via the --league option")
		}
		league, err := leagues.New(context.String("league"))
		if err != nil {
			return err
		}
		fmt.Println("Disciplines:")
		for k, v := range league.Disciplines() {
			fmt.Printf("- [%s] %s\n", k, v)
		}
		return nil
	},
}

var leagueChallengesCommand = cli.Command{
	Name:  "challenges",
	Usage: "List all challenges",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "league",
			Usage: "the id of the league",
		},
	},
	Action: func(context *cli.Context) error {
		if !context.IsSet("league") {
			return fmt.Errorf("Please specify the id of the league to used via the --league option")
		}
		league, err := leagues.New(context.String("league"))
		if err != nil {
			return err
		}
		fmt.Println("Challenges:")
		for k, v := range league.Challenges() {
			fmt.Printf("- [%s] %s\n", k, v)
		}
		return nil
	},
}

var leagueResultsCommand = cli.Command{
	Name:  "results",
	Usage: "Display results for a challenge",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "league",
			Usage: "the id of the league",
		},
		cli.StringFlag{
			Name:  "discipline",
			Usage: "the discipline of the challenge",
		},
		cli.StringFlag{
			Name:  "level",
			Usage: "the level of the challenge",
		},
		cli.StringFlag{
			Name:  "challenge",
			Usage: "the challenge",
		},
	},
	Action: func(context *cli.Context) error {
		if !context.IsSet("league") {
			return fmt.Errorf("Please specify the id of the league to used via the --league option")
		}
		if !context.IsSet("discipline") {
			return fmt.Errorf("Please specify the id of the discipline to used via the --discipline option")
		}
		if !context.IsSet("level") {
			return fmt.Errorf("Please specify the id of the level to used via the --level option")
		}
		if !context.IsSet("challenge") {
			return fmt.Errorf("Please specify the id of the challenge to used via the --challenge option")
		}
		league, err := leagues.New(context.String("league"))
		if err != nil {
			return err
		}
		league.Display(context.String("challenge"), context.String("discipline"), context.String("level"))
		return nil
	},
}
