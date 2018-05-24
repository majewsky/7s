(function() {

  var setCurrentSlideInUI = function(idx) {
    document.querySelector("iframe#prev").contentDocument.location.href = '/slides/' + (idx - 1);
    document.querySelector("iframe#curr").contentDocument.location.href = '/slides/' + (idx + 0);
    document.querySelector("iframe#next").contentDocument.location.href = '/slides/' + (idx + 1);

    window.currentSlide = idx;
  };

  setCurrentSlideInUI(window.currentSlide);

  var setCurrentSlideOnServer = function(idx) {
    var req = new XMLHttpRequest();
    req.addEventListener("load", function(e) {
      console.log(e.target.responseText);
      var newIdx = JSON.parse(e.target.responseText)["current_slide"];
      if (newIdx !== undefined) {
        setCurrentSlideInUI(newIdx);
      }
    });

    req.open("POST", document.location.pathname + '?set-slide=' + idx);
    req.send();
  };

  document.querySelector("button#btn-prev").onclick = function() {
    setCurrentSlideOnServer(window.currentSlide - 1);
  };
  document.querySelector("button#btn-next").onclick = function() {
    setCurrentSlideOnServer(window.currentSlide + 1);
  };

}());
