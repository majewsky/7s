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
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
)

//Server represents the state of the HTTP server.
type Server struct {
	Deck         *Deck
	PresenterKey string
}

//Run starts the HTTP server in this goroutine. It does not return, except in
//error cases.
func (s *Server) Run(listenAddress string) error {
	r := mux.NewRouter().StrictSlash(true)

	rGET := r.Methods("GET").Subrouter()
	rGET.HandleFunc("/slides/{idx:[0-9]+}", s.reqGetSlide)
	rGET.HandleFunc("/assets/{path:.+}", s.reqGetAsset)

	return http.ListenAndServe(listenAddress, r)
}

func (s *Server) reqGetSlide(w http.ResponseWriter, r *http.Request) {
	idx, err := strconv.ParseUint(mux.Vars(r)["idx"], 10, 64)
	if err != nil || idx >= uint64(len(s.Deck.SlidesHTML)) {
		http.NotFound(w, r)
		return
	}

	Page{
		Content: s.Deck.SlidesHTML[idx],
		Title:   fmt.Sprintf("Slide %d", idx),
		BaseURL: "/assets/",
	}.WriteTo(w)
}

func (s *Server) reqGetAsset(w http.ResponseWriter, r *http.Request) {
	//only expose assets that are referenced by a slide
	path := mux.Vars(r)["path"]
	if !s.Deck.AssetPaths[path] {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, filepath.Join(s.Deck.BaseDirectory, path))
}
