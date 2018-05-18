/*******************************************************************************
*
* Copyright 2018 Stefan Majewsky <majewsky@gmx.net>
*
* This program is free software: you can redistribute it and/or modify it under
* the terms of the GNU General Public License as published by the Free Software
* Foundation, either version 3 of the License, or (at your option) any later
* version.
*
* This program is distributed in the hope that it will be useful, but WITHOUT ANY
* WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
* A PARTICULAR PURPOSE. See the GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License along with
* this program. If not, see <http://www.gnu.org/licenses/>.
*
*******************************************************************************/

package main

import (
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"github.com/golang-commonmark/markdown"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

//Deck represents the entire slide deck.
type Deck struct {
	//BaseDirectory is the path to the directory that contains the slide deck
	//file.
	BaseDirectory string
	//SlidesHTML contains the generated HTML for each individual slide.
	SlidesHTML []string
	//AssetPaths contains the paths of all assets (as of now, only images)
	//referenced by the slides, which need to be served by 7s from the file
	//system. Paths are relative to the BaseDirectory.
	AssetPaths map[string]bool
}

//NewDeck parses the slide deck file at the given path.
func NewDeck(path string) (*Deck, error) {
	deckBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	d := &Deck{
		BaseDirectory: filepath.Dir(path),
		AssetPaths:    make(map[string]bool),
	}

	//split deck into separate Markdown documents for each slide, compile into HTML
	md := markdown.New(markdown.HTML(true))
	for _, slideStr := range slideSeparatorRx.Split(string(deckBytes), -1) {
		if strings.IndexFunc(slideStr, isNotSpace) != -1 {
			d.SlidesHTML = append(d.SlidesHTML, md.RenderToString([]byte(slideStr)))
		}
	}

	//find URLs of assets
	for _, slideHTML := range d.SlidesHTML {
		scanAssets(slideHTML, d.AssetPaths)
	}

	return d, nil
}

var slideSeparatorRx = regexp.MustCompile(`(?m)^\s*---\s*$`)

func isNotSpace(r rune) bool {
	return !unicode.IsSpace(r)
}

func scanAssets(slideHTML string, assetPaths map[string]bool) {
	tokenizer := html.NewTokenizer(strings.NewReader(slideHTML))
	for {
		switch tokenizer.Next() {
		case html.ErrorToken:
			//end of document
			return
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.DataAtom == atom.Img {
				//found an <img> tag -- find its src attribute
				for _, attr := range token.Attr {
					if attr.Key == "src" {
						assetPaths[filepath.Clean(attr.Val)] = true
						break
					}
				}
			}
		}
	}
}
