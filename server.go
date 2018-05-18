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
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/gorilla/mux"
)

//Server represents the state of the HTTP server.
type Server struct {
	Deck         *Deck
	PresenterKey string
	currentIdx   uint64
	currentMutex *sync.RWMutex
}

//Run starts the HTTP server in this goroutine. It does not return, except in
//error cases.
func (s *Server) Run(listenAddress string) error {
	if s.currentMutex == nil {
		s.currentIdx = 1
		s.currentMutex = &sync.RWMutex{}
	}

	r := mux.NewRouter().StrictSlash(true)

	rGET := r.Methods("GET").Subrouter()
	rGET.HandleFunc("/assets/7s.css", s.hardcodedAsset("assets/7s.css", []byte(cssSlides)))
	rGET.HandleFunc("/assets/7s.js", s.hardcodedAsset("assets/7s.js", []byte(jsSlides)))
	rGET.HandleFunc("/presenter/7s.css", s.hardcodedAsset("presenter/7s.css", []byte(cssPresenter)))
	rGET.HandleFunc("/presenter/7s.js", s.hardcodedAsset("presenter/7s.js", []byte(jsPresenter)))

	rGET.HandleFunc("/slides/{idx:[0-9]+}", s.reqGetSlide)
	rGET.HandleFunc("/assets/{path:.+}", s.reqGetAsset)
	rGET.HandleFunc("/"+s.PresenterKey, s.reqGetPresenter)

	return http.ListenAndServe(listenAddress, r)
}

func (s *Server) reqGetSlide(w http.ResponseWriter, r *http.Request) {
	idx, err := strconv.ParseUint(mux.Vars(r)["idx"], 10, 64)
	if err != nil {
		idx = 0
	}

	p := Page{
		Content:    "<p>No such slide</p>",
		CSSClasses: "no-slide",
		Title:      fmt.Sprintf("Slide %d", idx),
		BaseURL:    "/assets/",
	}
	if idx > 0 && idx <= uint64(len(s.Deck.SlidesHTML)) {
		p.Content = s.Deck.SlidesHTML[idx-1]
		p.CSSClasses = "slide"
	}
	p.WriteTo(w)
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

func (s *Server) hardcodedAsset(path string, contents []byte) http.HandlerFunc {
	//This handler serves {slides,presenter}/7s.{js,css}. It usually uses the
	//builtin files from static.go, but if an override file with the same path is
	//present in the s.Deck.BaseDirectory, that one is served instead. This
	//behavior is intended for development, but may also be useful if users want
	//to apply custom styles to their slides.
	fsPath := filepath.Join(s.Deck.BaseDirectory, path)
	return func(w http.ResponseWriter, r *http.Request) {
		fi, err := os.Stat(fsPath)
		if err == nil && fi.Mode().IsRegular() {
			http.ServeFile(w, r, fsPath)
			return
		}
		if strings.HasSuffix(path, "css") {
			w.Header().Set("Content-Type", "text/css")
		} else {
			w.Header().Set("Content-Type", "text/javascript")
		}
		w.WriteHeader(200)
		w.Write(contents)
	}
}

func (s *Server) reqGetPresenter(w http.ResponseWriter, r *http.Request) {
	Page{
		Content:    templPresenter,
		CSSClasses: "presenter",
		Title:      "7s Presenter",
		BaseURL:    "/presenter/", //load different CSS/JS here
	}.WriteTo(w)
}
