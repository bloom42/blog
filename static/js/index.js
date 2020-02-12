// document ready
function rot13(s) {
    return (s ? s : this).split('').map(function(_){
      if (!_.match(/[A-Za-z]/)) return _;
      c = Math.floor(_.charCodeAt(0) / 97);
      k = (_.toLowerCase().charCodeAt(0) - 83) % 26 || 26;
      return String.fromCharCode(k + ((c == 0) ? 64 : 96));
    }).join('');
}
(function ($) {

  var previousScroll = 20;
  var minimumScroll = 71;
      // scroll functions
      $(window).scroll(function(e) {

          // add/remove class to navbar when scrolling to hide/show
          var scroll = $(window).scrollTop();
          if (scroll >= previousScroll && scroll >= minimumScroll) {
              $('.navbar').addClass("navbar-hide");
          } else if (scroll < previousScroll) {
              $('.navbar').removeClass("navbar-hide");
          }
          previousScroll = scroll;
      });

    // deofuscate emails
    var emails = document.getElementsByClassName("obfuscated-email");
    if (emails) {
      for (var i = 0; i < emails.length; i += 1) {
        emails[i].innerHTML = rot13(`<n uers="znvygb:uryyb@oybbz.fu">uryyb@oybbz.fu</n>`);
      }
    }

})(jQuery);
