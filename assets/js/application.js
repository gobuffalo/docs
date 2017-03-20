require('expose-loader?$!expose-loader?jQuery!jquery');
require("bootstrap/dist/js/bootstrap.js");

$(function() {
  activateSideNav();
  $('.highlight pre').each(function(i, block) {
    var html = block.innerHTML;
    html = html.replace(/\t/g, "  ");
    block.innerHTML = html;
    hljs.highlightBlock(block);
  });

  $("img[alt='Buffalo Logo']").closest("p").css("text-align", "center");

  $(window).on("hashchange", activateSideNav);

  buildSideNav();
});

function buildSideNav() {
  loc = window.location;
  var path = loc.pathname;
  var items = [];
  $(".main a[name]").each(function(_, a) {
    a = $(a);
    if (a.attr("title")) {
      items.push(`<li>> <a href="${path}#${a.attr('name')}">${a.attr("title")}</a></li>`);
    }
  })
  $("#topics").addClass("list-unstyled");
  $("#topics").append(items);
}

function activateSideNav() {
  loc = window.location;
  var path = loc.pathname;
  $(".nav-sidebar li").removeClass("active")
  $(`.nav a[href="${path}"]`).closest("li").addClass("active");
}
