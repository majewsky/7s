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
	"net/http"
	"strings"
)

const templOuter = `<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<base href="%BASE%">
	<link rel="stylesheet" type="text/css" href="7s.css">
	<title>%TITLE%</title>
</head>
<body class="%BODYCLASS%">
	%CONTENT%
	<script type="text/javascript" src="7s.js"></script>
</body>`

//Page represents a full HTML page.
type Page struct {
	Content    string
	Title      string
	CSSClasses string
	BaseURL    string
}

//WriteTo writes the page.
func (p Page) WriteTo(w http.ResponseWriter) {
	if p.BaseURL == "" {
		p.BaseURL = "/"
	}

	out := templOuter
	out = strings.Replace(out, "%CONTENT%", p.Content, 1)
	out = strings.Replace(out, "%BODYCLASS%", p.CSSClasses, 1)
	out = strings.Replace(out, "%TITLE%", p.Title, 1)
	out = strings.Replace(out, "%BASE%", p.BaseURL, 1)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte(out))
}

////////////////////////////////////////////////////////////////////////////////

const templPresenter = `
	<iframe id="prev" src="/slides/0"></iframe>
	<iframe id="curr" src="/slides/0"></iframe>
	<iframe id="next" src="/slides/0"></iframe>
	<button id="btn-prev">« Prev</button>
	<button id="btn-next">Next »</button>

	<script type="text/javascript">window.currentSlide = %SLIDENUM%;</script>
`

const templAudience = `
	<p id="noscript">Please enable JavaScript for this domain.</p>
	<script type="text/javascript">window.currentSlide = %SLIDENUM%;</script>
`
