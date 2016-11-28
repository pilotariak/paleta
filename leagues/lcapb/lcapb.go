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

package lcapb

import (
	// "fmt"

	"github.com/Sirupsen/logrus"

	"github.com/pilotariak/paleta/leagues"
)

const (
	uri = "http://lcapb.euskalpilota.fr/resultats.php"
)

var (
	current = "20170501"

	disciplines = map[string]string{
		"2":   "Trinquet / P.G. Pleine Masculin",
		"3":   "Trinquet / P.G. Creuse Masculin",
		"4":   "Trinquet / P.G. Pleine Feminine",
		"5":   "Trinquet / P.G. Creuse Feminine",
		"13":  "Place Libre / Grand Chistera",
		"16":  "Place Libre / P.G. Pleine Masculin",
		"26":  "Mur à Gauche / P.G. Pleine Masculin",
		"27":  "Mur à Gauche / P.G. Pleine Feminine",
		"28":  "Mur à Gauche / P.G. Creuse Masculin Individuel",
		"126": "Mur A gauche / P.G. Pleine Masculin Barrages",
		"501": "Place Libre / P.G Pleine Feminine",
	}

	levels = map[string]string{
		"1":  "1ère Série",
		"2":  "2ème Série",
		"3":  "3ème Série",
		"4":  "Seniors",
		"6":  "Cadets",
		"7":  "Minimes",
		"8":  "Benjamins",
		"9":  "Poussins",
		"51": "Senoir Individuel",
	}
)

func init() {
	leagues.RegisterLeague("lcapb", newLCAPBLeague)
}

type lcapLeague struct {
	Website string
}

func newLCAPBLeague() (leagues.League, error) {
	return &lcapLeague{
		Website: "http://www.lcapb.net/",
	}, nil
}

func (l *lcapLeague) Levels() map[string]string {
	return levels
}

func (l *lcapLeague) Disciplines() map[string]string {
	return disciplines
}

func (l *lcapLeague) Display(disciplineID string, levelID string) error {

	logrus.Debugf("[lcapb] Search results for %s %s", disciplineID, levelID)
	return leagues.Display(uri, disciplineID, levelID, current)
}
