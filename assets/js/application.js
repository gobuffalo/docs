require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.js");
require("./prism.js");
require("./theme.js");
require("expose-loader?Clipboard!./clipboard.min.js");

let buildSideNav = () => {
    let loc = window.location;
    let path = loc.pathname;
    let sb = $(`a[href="${path}"]`);
    sb.closest("ul.sidenav").addClass("open");

    let items = [];
    $(".main-content a[name]").each((_, a) => {
        let $a = $(a);
        if ($a.attr("title")) {
            let name = $a.attr("name");
            let title = $a.attr("title");
            items.push(`<li><a href="${path}#${name}">${title}</a></li>`);
        }
    });
    if (items.length > 0) {
        let ul = $("<ul class='sub-nav'>").append(items);
        sb.append(ul);
        sb.addClass("active");
        sb.addClass("has-child");
        sb.addClass("open");
    }
};

let activateSideNav = () => {
    let loc = window.location;
    let path = loc.pathname;
    $(".sidebar li").removeClass("active");
    let item = $(`.sidebar a[href="${path}"]`);
    item.closest("li").addClass("active");
};

$(() => {

    activateSideNav();
    buildSideNav();

    $(".code-tabs .window-content").each((_, wc) => {
        $(wc).find("pre").first().show();
    });

    $("img[title=screenshot]").addClass("img-shadow img-responsive center-block img-rounded");

    $("p:empty").remove();
    $("code:empty").remove();
});