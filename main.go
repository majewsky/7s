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
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
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

	presenterKey, err := generatePresenterKey()
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not generate secret key for presenter: %s\n", err.Error())
	}
	deck, err := NewDeck(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	log.Printf("serving %d slides with %d assets\n",
		len(deck.SlidesHTML), len(deck.AssetPaths))
	hostPort := os.Args[2]
	if strings.HasPrefix(hostPort, ":") {
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "localhost"
		}
		hostPort = hostname + hostPort
	}
	log.Printf("private URL for presenter is http://%s/%s\n", hostPort, presenterKey)
	log.Printf("public URL for audience is http://%s\n", hostPort)
	log.Println("NOTE: URLs may differ when users reach 7s through reverse proxies or NAT")
}

func generatePresenterKey() (string, error) {
	var buf [16]byte
	_, err := rand.Read(buf[:])
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(buf[:]), nil
}
