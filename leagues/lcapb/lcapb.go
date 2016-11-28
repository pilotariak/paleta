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
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"

	"github.com/pilotariak/paleta/leagues"
)

const (
	uri = "http://lcapb.euskalpilota.fr/resultats.php"
)

var (
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
	return &lcapLeague{}, nil
}

func (l *lcapLeague) Levels() map[string]string {
	return levels
}

func (l *lcapLeague) Disciplines() map[string]string {
	return disciplines
}

func fetch(disciplineID string, levelID string) ([]byte, error) {
	data := url.Values{}
	data.Add("InSel", "")
	data.Add("InCompet", "20170501")
	data.Add("InSpec", disciplineID)
	data.Add("InVille", "0")
	data.Add("InClub", "0")
	data.Add("InDate", "")
	data.Add("InDatef", "")
	data.Add("InCat", levelID)
	data.Add("InPhase", "0")
	data.Add("InPoule", "0")
	data.Add("InGroupe", "0")
	data.Add("InVoir", "Voir les résultats")
	u, _ := url.ParseRequestURI(uri)
	urlStr := fmt.Sprintf("%v", u)

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	resp, err := client.Do(r)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Http request to %s failed: %s", r.URL, err.Error())
	}
	fmt.Println(resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return nil, fmt.Errorf("errorination happened reading the body: %s", err.Error())
	}
	return body, nil
}

func (l *lcapLeague) Display(disciplineID string, levelID string) error {
	body, err := fetch(disciplineID, levelID)
	if err != nil {
		return err
	}
	z := html.NewTokenizer(strings.NewReader(string(body)))

	content := []string{"", "", "", "", ""}
	i := -1
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Club 1", "Club 2", "Score", "Commentaire"})
	table.SetRowLine(true)
	table.SetAutoWrapText(false)
	for {
		// token type
		tokenType := z.Next()
		if tokenType == html.ErrorToken {
			break
		}
		// token := z.Token()
		switch tokenType {
		case html.StartTagToken: // <tag>
			t := z.Token()
			if t.Data == "tr" {
				i = -1

			} else if t.Data == "td" {
				inner := z.Next()
				if inner == html.TextToken {
					if len(t.Attr) > 0 {
						if t.Attr[0].Val == "L0" { // Text to extract
							text := (string)(z.Text())
							value := strings.TrimSpace(text)
							if len(value) > 0 {
								i = i + 1
								// fmt.Printf("%d Attr::::::::::: %s :: %s\n", i, value, t.Attr)
								content[i] = value
							}
						} else if t.Attr[0].Val == "mTitreSmall" {
							text := (string)(z.Text())
							value := strings.TrimSpace(text)
							if len(value) > 0 {
								i = i + 1
								// fmt.Printf("%d Attr::::::::::: %s :: %s\n", i, value, t.Attr)
								content[i] = value
							}
						}
					}
				}

			} else if t.Data == "li" {
				inner := z.Next()
				if inner == html.TextToken {
					text := (string)(z.Text())
					value := strings.TrimSpace(text)
					// fmt.Printf("%s\n%s", content[i], value)
					content[i] = fmt.Sprintf("%s\n%s", content[i], value)
				}

			}
		case html.TextToken: // text between start and end tag
		case html.EndTagToken: // </tag>
			t := z.Token()
			if t.Data == "tr" {
				if len(content[0]) > 0 {
					// fmt.Printf("==> %d\n", len(content))
					for rank, elem := range content {
						fmt.Printf("%d = %s\n", rank, elem)
					}
					table.Append(content)
					content = []string{"", "", "", "", ""}
				}
			}

		case html.SelfClosingTagToken: // <tag/>
		}
	}

	table.Render()
	return nil
}
