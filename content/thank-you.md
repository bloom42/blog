+++
title = "Thank you!"
type = "page"
date = 2019-02-05T01:42:42+02:00
+++

<br />

# Thank you for subscribing!

<div class="text-left">

<br />

<div id="osw-alert"></div>

<div class="spinner-border text-primary" role="status" id="osw-loader">
  <span class="sr-only">Loading...</span>
</div>

The next issue will be coming soon, but in the mean time you can check out the **[last issues](/)**.

Also, just so you know, you can always click reply to any issue to give feedback or ask questions.

</div>


<script type="application/javascript">
$(document).ready(function() {
  var urlParams = new URLSearchParams(window.location.search);
  var token = urlParams.get('token');
  var contactId = urlParams.get('contact');
  if (token == null || contactId == null) {
    hideLoader();
    displayError('Error confirming your subscription. Please click again on the link provided in the email.');
    return;
  }

  var query = 'mutation($input: VerifyContactInput!) { verifyContact(input: $input) }';
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
    displaySuccess('You are now subscribed to <b>Open Source Weekly</b>');
  })
  .fail(function() {
    displayError('Error confirming your subscription. Please click again on the link provided in the email.');
  })
  .always(function() {
    hideLoader();
  })
});
</script>
