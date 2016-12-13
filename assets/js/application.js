$(function() {
  var path = window.location.pathname;
  $(".nav a[href='" + path + "']").closest("li").addClass("active");

  $('.highlight pre').each(function(i, block) {
    hljs.highlightBlock(block);
  });
});
