+++
title = "About"
type = "page"
date = 2019-02-05T01:42:42+02:00
+++

# About
<!--
We believe that entropy (which can be defined as a measurement of chaos, complexity) is inevitable, and equilibrium is death.
As product designers, developers, entrepreneurs, managers... it's our role to manage complexity.
We share on this blog our ideas and experience on how we tackle entropy and keep things simple. -->

Equilibrium is death but too much entropy (chaos) is fatal.
In this blog we explore how complexity and chaos infiltrate systems, organizations and organisms and lead to a slow and painful death.

<!-- Sometime, it's the lack of entropy that can be fatal: in cryptography or in lottery for example. Anyway -->

Entropy is an important part of our lives and if you don't control it, it will control you.

<br />

## Authors

From the creators (and friends) of [Bloom](https://bloom.sh). Their mission is to empower
people with open technologies ✌️

<br />

## Contact

You can contact the authors using one of the following way:
* Using email: <span id="obfuscated-email">Javascript is required to access email address</span>
* By saying <i>Hi</i> on Twitter: <a href="https://twitter.com/42bloom" target="_blank" rel="noopener">@42bloom</a>

<br />

## License

All the content on this website is licensed under the
<a rel="noopener" target="_blank" href="https://creativecommons.org/licenses/by-sa/4.0/" >CC BY-SA 4.0 License</a>.

<br />

## Source code

The source code of this website is available on GitLab:
<a href="{{< giturl >}}" target="_blank" rel="noopener">{{< giturl >}}</a>


<script type="text/javascript">
  window.addEventListener("load", function(){
    var email = document.getElementById("obfuscated-email");
    if (email) {
        email.innerHTML = rot13('<n uers="znvygb:uryyb@oybbz.fu">uryyb@oybbz.fu</n>');
    }
  });
  function rot13(s) {
    return (s ? s : this).split('').map(function(_){
      if (!_.match(/[A-Za-z]/)) return _;
      c = Math.floor(_.charCodeAt(0) / 97);
      k = (_.toLowerCase().charCodeAt(0) - 83) % 26 || 26;
      return String.fromCharCode(k + ((c == 0) ? 64 : 96));
    }).join('');
  }
</script>
