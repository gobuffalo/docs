document.addEventListener('DOMContentLoaded', () =>{
    document.querySelectorAll("h1, h2, h3, h4, h5, h6").forEach(h => {
        if (h.id == ""){
            return
        }

        let link = document.createElement("a");
        link.href = "#" + h.id;
        link.innerHTML = "#";
        link.classList.add("ml-2", "inline-block", "no-underline");

        h.innerHTML += link.outerHTML;
    })
});