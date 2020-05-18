// function rot13(s) {
//     return (s ? s : this).split('').map(function(_){
//       if (!_.match(/[A-Za-z]/)) return _;
//       c = Math.floor(_.charCodeAt(0) / 97);
//       k = (_.toLowerCase().charCodeAt(0) - 83) % 26 || 26;
//       return String.fromCharCode(k + ((c == 0) ? 64 : 96));
//     }).join('');
// }

// // document ready
// $(document).ready(function() {

//   var previousScroll = 20;
//   var minimumScroll = 71;
//   // scroll functions
//   $(window).scroll(function(e) {
//     // add/remove class to navbar when scrolling to hide/show
//     var scroll = $(window).scrollTop();
//     if (scroll >= previousScroll && scroll >= minimumScroll) {
//         $('.navbar').addClass("navbar-hide");
//     } else if (scroll < previousScroll) {
//         $('.navbar').removeClass("navbar-hide");
//     }
//     previousScroll = scroll;
//   });

//   // deofuscate emails
//   var emails = document.getElementsByClassName("obfuscated-email");
//   if (emails) {
//     for (var i = 0; i < emails.length; i += 1) {
//       emails[i].innerHTML = rot13(window.obfuscatedEmail);
//     }
//   }

//   // open all external links in new tab
//   $(document.links).filter(function() {
//     return this.hostname != window.location.hostname;
//   }).attr('target', '_blank').attr('rel', 'noopener');
// });
// Setup
window.apiBaseUrl = 'https://shibui.bloom42.com/api';
// window.apiBaseUrl = 'http://localhost:8000/api';
window.shibuiProjectId = 'ONUWIORPF5ZWQ2LCOVUS6UDSN5VGKY3UF4YQ';

// Functions
function rot13(s) {
    return (s ? s : this).split('').map(function(_){
      if (!_.match(/[A-Za-z]/)) return _;
      c = Math.floor(_.charCodeAt(0) / 97);
      k = (_.toLowerCase().charCodeAt(0) - 83) % 26 || 26;
      return String.fromCharCode(k + ((c == 0) ? 64 : 96));
    }).join('');
}

function displayError(message) {
  $('#osw-alert').html('<div class="alert alert-danger" role="alert">'
    + message +
  '</div>');
  $('#osw-alert').show();
}


function displaySuccess(message) {
  $('#osw-alert').html('<div class="alert alert-success" role="alert">'
    + message +
  '</div>');
  $('#osw-alert').show();
}

function hideAlert() {
  $('#osw-alert').hide();
}

function hideLoader() {
  $('#osw-loader').hide();
}

function showLoader() {
  $('#osw-loader').show();
}

function graphqlReq(query, variables) {
  var payload = {
    query: query,
    variables: variables,
  };
  var data = {
    data : JSON.stringify(payload),
    contentType : 'application/json',
    type : 'POST',
  };
  return $.ajax(window.apiBaseUrl+'/graphql', data);
}

// document ready
$(document).ready(function() {

  var previousScroll = 20;
  var minimumScroll = 71;
      // scroll functions
      $(window).scroll(function(e) {

          // add/remove class to navbar when scrolling to hide/show
          var scroll = $(window).scrollTop();
          if (scroll >= previousScroll && scroll >= minimumScroll) {
              $('nav').addClass("nav-hide");
          } else if (scroll < previousScroll) {
              $('nav').removeClass("nav-hide");
          }
          previousScroll = scroll;
      });

    // deofuscate emails
    var emails = document.getElementsByClassName("obfuscated-email");
    if (emails) {
      for (var i = 0; i < emails.length; i += 1) {
        emails[i].innerHTML = rot13(window.obfuscatedEmail);
      }
    }

    $(document.links).filter(function() {
      return this.hostname != window.location.hostname;
    }).attr('target', '_blank').attr('rel', 'noopener');
});
