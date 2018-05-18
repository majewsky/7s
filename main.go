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
	"fmt"
	"os"
)

const usageMessage = "usage: 7s <slide-deck-file> <listen-address>\n e.g. 7s /path/to/slides.7s.md :8080"

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--help" {
		fmt.Fprintln(os.Stdout, usageMessage)
		return
	}
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, usageMessage)
		os.Exit(1)
	}

	deck, err := NewDeck(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	//TODO DEBUG
	fmt.Fprintf(os.Stdout, "deck = %#v\n", deck)
}
