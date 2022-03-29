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

    fetch("https://api.rss2json.com/v1/api.json?rss_url=https://blog.gobuffalo.io/feed").then(response => {
        response.json().then(data => {
            let items = data.items.slice(0, 3);
            items.forEach(item => {
                container = document.querySelector("#blog-content");

                let desc = item.description.replace(/<img[^>]*>/g, "");
                desc = desc.replace(/<\/?[^>]+(>|$)/g, "");
                desc = desc.replace(/&nbsp;/g, "");
                desc = desc.replace(/&#8217;/g, "'");
                desc = desc.slice(0, 200) + "...";

                container.innerHTML += `
                <div class="">
                    <h4 class="text-2xl font-bold" >${item.title}</h4>
                    <p class="text-xs mb-3">${item.categories.join(", ")}</p>
                    <p class="text-left">${desc} <a class="underline" href="${item.link}">Read more</a></p>
                </div>
                `
            })


        })
    })
});