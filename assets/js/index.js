// document ready
$(document).ready(function() {

  // deofuscate emails
  var emails = document.getElementsByClassName("obfuscated-email");
  if (emails) {
    for (var i = 0; i < emails.length; i += 1) {
      emails[i].innerHTML = rot13(window.obfuscatedEmail);
    }
  }

  // open all links in new tab
  $(document.links).filter(function() {
    return this.hostname != window.location.hostname;
  }).attr('target', '_blank').attr('rel', 'noopener');
});
