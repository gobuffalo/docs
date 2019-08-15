import anchorJS from "anchor-js";
let anchors = new anchorJS();

// Auto-anchor titles
document.addEventListener("DOMContentLoaded", function(event) {
    anchors.add();
    buildToc();
    activateSideNav();
});

// Build the table of contents
var buildToc = () => {
    let loc = window.location;
    let path = loc.pathname.replace(/\/$/, "");
    let items = [];
    $(".main-content h2").each((_, a) => {
        let name = anchors.urlify(a.textContent);
        let title = a.textContent;
        items.push(`<li><a href="${path}#${name}">${title}</a></li>`);
    });
    if (items.length > 0) {
        let ul = $("<ul class=\"summary\">").append(items);
        $(".main-content h1:first").after(ul);
    }
};

// Activate entry in the menu
var activateSideNav = () => {
    let loc = window.location;
    let path = loc.pathname === "/" ? "/docs/overview" : loc.pathname;
    if (path[path.length - 1] !== "/") {
        path += "/";
    }
    $(".sidebar-root li").removeClass("active");
    let item = $(`.sidebar a[href="${path}"]`);
    item.parentsUntil("ul.sidebar-root", "li").addClass("active");
    let sn = item.parentsUntil("ul.sidebar-root");
    sn.addClass("open");
};