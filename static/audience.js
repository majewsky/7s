(function() {

  //remove noscript indicator
  var nsNode = document.getElementById("noscript");
  if (nsNode) {
    nsNode.parentNode.removeChild(nsNode);
  }

  //add <iframe> that displays the current slide
  var iframe = document.createElement("iframe");
  iframe.id = "curr";
  document.querySelector("body").appendChild(iframe);

  var setCurrentSlideInUI = function(idx) {
    document.querySelector("iframe#curr").contentDocument.location.href = '/slides/' + (idx + 0);
    window.currentSlide = idx;
  };
  setCurrentSlideInUI(window.currentSlide);

  //long-poll loop that advances through the slides as indicated by the server
  var poll;
  poll = function() {
    var req = new XMLHttpRequest();
    req.addEventListener("error", function() {
      console.log("GET /notify failed, restarting in 2 seconds...");
      setTimeout(poll, 2000);
    });
    req.addEventListener("abort", function() {
      console.log("GET /notify aborted! Giving up.");
    });
    req.addEventListener("load", function(e) {
      if (e.target.status !== 200) {
        console.log("GET /notify returned error status, restarting in 2 seconds...");
        setTimeout(poll, 2000);
        return;
      }
      var newIdx = JSON.parse(e.target.responseText)["current_slide"];
      if (newIdx !== undefined && newIdx !== window.currentSlide) {
        setCurrentSlideInUI(newIdx);
      }
      poll();
    });

    req.open("GET", "/notify?current=" + window.currentSlide);
    req.send();
  };
  poll();

}());
