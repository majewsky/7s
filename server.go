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
	rGET.HandleFunc("/assets/7s.css", s.hardcodedAsset("static/slides.css"))
	rGET.HandleFunc("/assets/7s.js", s.hardcodedAsset("static/slides.js"))
	rGET.HandleFunc("/presenter/7s.css", s.hardcodedAsset("static/presenter.css"))
	rGET.HandleFunc("/presenter/7s.js", s.hardcodedAsset("static/presenter.js"))

	rGET.HandleFunc("/slides/{idx:[0-9]+}", s.reqGetSlide)
	rGET.HandleFunc("/assets/{path:.+}", s.reqGetAsset)
	rGET.HandleFunc("/"+s.PresenterKey, s.reqGetPresenter)

	r.HandleFunc("/"+s.PresenterKey, s.reqPostPresenter).Methods("POST")

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

func (s *Server) hardcodedAsset(path string) http.HandlerFunc {
	//This handler serves {assets,presenter}/7s.{js,css}. It usually uses the
	//builtin files from static.go, but if the application is run from its Git
	//repo, the corresponding source files from static/ are served instead. This
	//behavior is intended for development, but may also be useful if users want
	//to apply custom styles to their slides.
	fsPath := filepath.Join(filepath.Dir(os.Args[0]), "..", path)
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
		contents, _ := Asset(path)
		w.Write(contents)
	}
}

func (s *Server) reqGetPresenter(w http.ResponseWriter, r *http.Request) {
	s.currentMutex.RLock()
	currentIdx := strconv.FormatUint(s.currentIdx, 10)
	s.currentMutex.RUnlock()
	Page{
		Content:    strings.Replace(templPresenter, "%SLIDENUM%", currentIdx, -1),
		CSSClasses: "presenter",
		Title:      "7s Presenter",
		BaseURL:    "/presenter/", //load different CSS/JS here
	}.WriteTo(w)
}

func (s *Server) reqPostPresenter(w http.ResponseWriter, r *http.Request) {
	//TODO
	w.WriteHeader(500)
	w.Write([]byte("TODO"))
}
