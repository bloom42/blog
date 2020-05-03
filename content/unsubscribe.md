+++
title = "Thank you!"
type = "page"
date = 2019-02-05T01:42:42+02:00
+++

<br />

# Unsubscribe


<br />

<div id="osw-alert"></div>

<div class="spinner-border text-primary" role="status" style="display: none;" id="osw-loader">
  <span class="sr-only">Loading...</span>
</div>

Are you Sure?

<a href="/"><button type="button" class="btn btn-primary">Cancel</button></a>
<button type="button" class="btn btn-danger" id="osw-unsubscribe-btn">Yes, unsubscribe me</button>


<script type="application/javascript">
function unsubscribe() {
  hideAlert();
  var urlParams = new URLSearchParams(window.location.search);
  var token = urlParams.get('token');
  var contactId = urlParams.get('contact');
  if (token == null || contactId == null) {
    hideLoader();
    displayError('Error unsuscribing. Please click again on the link provided in the email.');
    return;
  }

  var query = 'mutation($input: UnsubscribeContactInput!) { unsubscribeContact(input: $input) }';
  var variables = {
    input: {
      id: decodeURIComponent(contactId),
      token: token,
    },
  };
  graphqlReq(query, variables)
  .done(function(data) {
    if (data.errors && data.errors.length > 0) {
       displayError(data.errors[0].message);
       return ;
    }
    displaySuccess('You no longer will receive our emails. Have a good day!');
  })
  .fail(function() {
    displayError('Error unsuscribing. Please click again on the link provided in the email.');
  })
  .always(function() {
    hideLoader();
  })
}

$(document).ready(function() {
  hideLoader();
  $("#osw-unsubscribe-btn").click(function(e) {
    e.preventDefault();
    unsubscribe();
  });
});
</script>
