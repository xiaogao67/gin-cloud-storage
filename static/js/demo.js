/**
 * demo.js
 * https://coidea.website
 *
 * Licensed under the MIT license.
 * http://www.opensource.org/licenses/mit-license.php
 *
 * Copyright 2018, COIDEA
 * https://coidea.website
 */

$(function() {

  $('.password-holder .unlock').click(function(e) {
    var errorMessage = $('.error-message');
    var passwordDiv = $(this).parent();

    if(passwordDiv.children('input').val()) {

      var tl = new TimelineMax();

      $("#download").submit()
      // tl.fromTo(passwordDiv, 0.3, {x:-1}, { x:1, ease:RoughEase.ease.config({ strength:8, points:40, template:Linear.easeNone, randomize:false }) , clearProps:"x" })
      //   .to($('body'), 0.15, { backgroundColor: '#E74C3C' })
      //   .to(errorMessage, 0.15, { autoAlpha: 1, y: -16 }, "-=0.15")
      //   .to(passwordDiv, 0.15, { className: "+=false" }, "-=0.15")
      //   .to(passwordDiv, 0.15, { className: "-=false" }, "+=2.5")
      //   .to($('body'), 0.15, { backgroundColor: '#EDF0F9' }, "+=0.65")
      //   .to(errorMessage, 0.15, { autoAlpha: 0, y: 0 }, "-=0.15");

    }
  })

});
