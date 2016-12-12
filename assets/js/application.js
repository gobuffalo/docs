$(function() {
  var path = window.location.pathname;
  $(".nav a[href='"+path+"']").closest("li").addClass("active");

  hljs.initHighlighting();
});
