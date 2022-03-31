require("expose-loader?$!expose-loader?jQuery!jquery");
require("./theme.js");
require("expose-loader?Clipboard!./clipboard.min.js");

import "bootstrap/dist/js/bootstrap.bundle"
//import highlighter from 'highlight.js'


$(() => {
  $(".code-tabs .window-content").each((_, wc) => {
    $(wc).find("pre").first().show();
  });

  $("img[title=screenshot]").addClass(
    "img-shadow img-responsive center-block img-rounded"
  );

  $(".codetabs").each((_, ct) => {
    let el = $(ct);
    let ul = el.find("ul:first-child");
    ul.addClass("nav nav-tabs");
    let tc = el.find(".tab-content");
    let id = el.attr("id");
    let blocks = el.find(".tab-content .highlight");
    blocks.each((i, b) => {
      let lid = `${id}-${i}`;
      let block = $(b);
      let name = block.text().split("\n")[0];
      name = name.toString();

      try {
        name = name.replace("Copy// ", "");

        let act = "";
        if (i === 0) {
          act = "active";
        }
        ul.append(
          `<li role="presentation" class="nav-item ${act}"><a href="#${lid}" role="tab" data-toggle="tab" class="nav-link ${act}">${name}</a></li>`
        );
        tc.append(
          $(
            `<div role="tabpanel" class="tab-pane ${act}" id="${lid}"></div>`
          ).append(block)
        );
        blocks.remove(block);
      } catch (err) {
        if (window.console) {
          console.log("err:", err);
        }
      }
    });
  });
});


$(() => {
  //Handle language switch
  $("#language").on("change", (e) => {
    $(e.target).
      closest("form").
      submit();
  });

  $("body").on("hidden.bs.modal", (e) => {
    var $iframes = $(e.target).find("iframe");
    $iframes.each((index, iframe) => {
      $(iframe).attr("src", $(iframe).attr("src"));
    });
  });

  // $('.highlight pre').each(function(_, block) {
  //   highlighter.highlightBlock(block);
  // });

  $(".faq h6").on("click", (e) => {
    let a = $(e.currentTarget).find("a[name]");
    let hash = a.attr("name");
    window.location.hash = hash;
  });

  let hash = window.location.hash;
  if (hash !== "") {
    if (hash.charAt(0) === "#") {
      hash = hash.slice(1);
    }
    $(`.faq h6 a[name="${hash}"]`).click();
  }
});

