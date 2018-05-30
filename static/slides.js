(function() {

  //apply special styling for slides with only one image on them (i.e. <body>
  //contains only one <p> and that <p> contains only one <img>)
  var nodes1 = document.querySelectorAll("body > :not(script)");
  if (nodes1.length == 1 && nodes1[0].tagName === "P") {
    var nodes2 = document.querySelectorAll("body > p > *");
    if (nodes2.length == 1 && nodes2[0].tagName === "IMG") {
      document.body.classList.add("only-one-image");
      //do not install the resize handler, we can handle this case without JS
      return;
    }
  }

  var resizeHandler = function() {
    var currentWidth = 50;
    var tooSmallWidth = undefined;
    var tooLargeWidth = undefined;
    var desiredHeight = window.innerHeight;
    var tooSmallHeight = undefined;
    var tooLargeHeight = undefined;

    // console.clear();
    // console.log("desired height = " + desiredHeight);
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
  window.onload = resizeHandler; //resize images again when all images loaded
  window.onresize = resizeHandler;

})();

