(function() {

  var setCurrentSlideInUI = function(idx) {
    document.querySelector("iframe#prev").contentDocument.location.href = '/slides/' + (idx - 1);
    document.querySelector("iframe#curr").contentDocument.location.href = '/slides/' + (idx + 0);
    document.querySelector("iframe#next").contentDocument.location.href = '/slides/' + (idx + 1);
  };

  setCurrentSlideInUI(window.currentSlide);

}());
