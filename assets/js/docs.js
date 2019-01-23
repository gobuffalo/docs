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
    $(".sidebar li").removeClass("active");
    let item = $(`.sidebar a[href="${path}"], .sidebar a[href="${path.substr(-1)}"]`);
    item.closest("li").addClass("active");
    let sn = item.closest("ul.sidenav");
    sn.addClass("open");
    sn.prev().addClass("open");
};