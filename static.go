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

const cssSlides = `
  html {
    font-size: 2vw;
    min-height: 100vh;
    overflow-y: hidden;
  }

  html, body, h1, h2, h3, h4, h5, h6, p {
    margin: 0;
    padding: 0;
    font-family: sans-serif;
  }

  body {
    padding: 1rem;
    padding-top: 0;
  }

  h1, h2, h3, h4, h5, h6 {
    margin: 0.5rem 0;
  }

  h1 {
    font-size: 3rem;
    line-height: 4rem;
  }
  h2 {
    font-size: 2.5rem;
    line-height: 3.5rem;
  }
  h3 {
    font-size: 2rem;
    line-height: 3rem;
  }
  h4 {
    font-size: 1.5rem;
    line-height: 2.5rem;
  }

  p, ul, ol, li {
    font-size: 1rem;
    line-height: 1.5rem;
  }

  p > img {
    display: block;
    margin: 0 auto;
  }

  body.no-slide {
    width: 100vw;
    height: 100vh;
    background: black;
    padding: 0;
  }
  body.no-slide p {
    line-height: 100vh;
    color: white;
    font-size: 4rem;
    text-align: center;
  }
`

const jsSlides = `
  (function() {

    var resizeHandler = function() {
      var currentWidth = 50;
      var tooSmallWidth = undefined;
      var tooLargeWidth = undefined;
      var desiredHeight = window.innerHeight;
      var tooSmallHeight = undefined;
      var tooLargeHeight = undefined;

      console.clear();
      console.log("desired height = " + desiredHeight);
      for (var iteration = 0; iteration < 50; iteration++) {
        //apply currentWidth
        var imageNodes = document.querySelectorAll("p > img");
        Array.prototype.forEach.call(imageNodes, function(imageNode) {
          imageNode.style.width = currentWidth + "vw";
        });

        //what height does that produce?
        var currentHeight = document.body.scrollHeight;
        if (currentHeight == desiredHeight) {
          break; //very unlikely
        }
        if (currentHeight < desiredHeight) {
          tooSmallWidth = currentWidth;
          tooSmallHeight = currentHeight;
          if (currentWidth >= 95) {
            //more than about 95vh does not make sense; the image would overflow on the right
            Array.prototype.forEach.call(imageNodes, function(imageNode) {
              imageNode.style.width = "95vw";
            });
            console.log("bound by width");
            break;
          }
        } else {
          tooLargeWidth = currentWidth;
          tooLargeHeight = currentHeight;
        }

        //do we have to explore more directions?
        if (tooSmallHeight === undefined) {
          currentWidth -= 20;
          if (currentWidth < 0) {
            //aspect ratio is too weird -> bail out
            Array.prototype.forEach.call(imageNodes, function(imageNode) {
              imageNode.style.width = "50vw";
            });
            console.log("bailing out");
            break;
          }
        } else if (tooLargeHeight == undefined) {
          currentWidth += 20;
        } else {
          //we have both a tooSmallWidth and a tooLargeWidth,
          //so bisect towards a good value
          currentWidth = (tooSmallWidth + tooLargeWidth) / 2;
        }

        console.log("(" + tooSmallWidth + " -> " + tooSmallHeight + ") - " + currentWidth + " - (" + tooLargeWidth + " -> " + tooLargeHeight + ")");

        //once converged, stop iterating
        if (tooSmallWidth / tooLargeWidth > 0.99) {
          console.log("breaking because converged");
          break;
        }
      }
    };

    resizeHandler();
    window.onresize = resizeHandler;

  })();
`

const cssPresenter = ``
const jsPresenter = ``
