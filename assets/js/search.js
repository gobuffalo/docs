
document.addEventListener('DOMContentLoaded', ()=>{  
  document.querySelectorAll(".search-button").forEach(element => {
    element.addEventListener('click', showSearch);
  })
    
  document.getElementById("search-overlay").addEventListener('click', hideSearch);
  document.getElementById("close-search").addEventListener('click', hideSearch);

  document.getElementById("search-input").addEventListener('keyup', searchTerm);
  document.addEventListener('keydown', handleKeyCodes)
});

function showSearch() {
  if (!document.querySelector("header").classList.contains("z-10")) {
    return
  }

  document.querySelector("header").classList.toggle("z-10");
  document.getElementById("search-palette").classList.toggle("hidden");
  document.getElementById("search-input").focus();

  
}

function hideSearch() {
  if (document.querySelector("header").classList.contains("z-10")) {
    return
  }

  document.querySelector("header").classList.toggle("z-10");
  document.getElementById("search-palette").classList.toggle("hidden");
}

// handleShow and hide search
function handleKeyCodes(e){
  if (e.keyCode === 27) {
    hideSearch();
  }

  if (e.keyCode >= 65 && e.keyCode <= 90) {
    let char = (e.metaKey ? '⌘-' : '') + String.fromCharCode(e.keyCode)
    if (char == "⌘-K") {
      showSearch();
    }
  }
}

function searchTerm(e) {
  let term = e.target.value;
  search(term)
}

let summaryInclude = 60;
var fuseOptions = {
  shouldSort: true,
  includeMatches: true,
  threshold: 0.0,
  tokenize:true,
  location: 0,
  distance: 100,
  maxPatternLength: 32,
  minMatchCharLength: 2,
  keys: [
    {name:"title",weight:0.8},
    {name:"contents",weight:0.5},
    {name:"tags",weight:0.3},
    {name:"categories",weight:0.3}
  ],
};

let tm = null
function search(searchQuery){
    if (searchQuery.length == 0) {
      document.getElementById("search-no-results").classList.add("hidden");
      document.getElementById("search-results").classList.add("hidden");
      document.getElementById("search-quick-actions").classList.remove("hidden");

      return
    }

    if (tm != null) {
      clearTimeout(tm)
    }

    tm = setTimeout(() => {
      fetch('/index.json').then(response => {
        response.json().then(data => {  
            var fuse = new Fuse(data, fuseOptions);
            var result = fuse.search(searchQuery);
            
            // hide the quick actions
            document.getElementById("search-quick-actions").classList.add("hidden");

            if (result.length == 0) {
              document.getElementById("search-no-results").classList.remove("hidden");
              document.getElementById("search-results").classList.add("hidden");
              return
            }

            document.getElementById("search-no-results").classList.add("hidden");
            document.getElementById("search-results").classList.remove("hidden");

            
            populateResults(result);
        });
      });
    }, 200);
}

function populateResults(result){
    let template = document.getElementById('search-result-template').innerHTML;
    document.getElementById('search-results').innerHTML = "";

    result.forEach((value, index) => {
      if (value.item.title == null) {
        return
      }

      var output = render(template,{
          key: index,
          title: value.item.title,
          link: value.item.permalink,
          tags: value.item.tags,
          categories: value.item.categories,
          snippet: value.item.contents
      });

      document.getElementById('search-results').innerHTML += output;
    });
}

function param(name) {
    return decodeURIComponent((location.search.split(name + '=')[1] || '').split('&')[0]).replace(/\+/g, ' ');
}

function render(templateString, data) {
  var conditionalMatches,conditionalPattern,copy;
  conditionalPattern = /\$\{\s*isset ([a-zA-Z]*) \s*\}(.*)\$\{\s*end\s*}/g;
  //since loop below depends on re.lastInxdex, we use a copy to capture any manipulations whilst inside the loop
  copy = templateString;
  while ((conditionalMatches = conditionalPattern.exec(templateString)) !== null) {
    if(data[conditionalMatches[1]]){
      //valid key, remove conditionals, leave contents.
      copy = copy.replace(conditionalMatches[0],conditionalMatches[2]);
    }else{
      //not valid, remove entire section
      copy = copy.replace(conditionalMatches[0],'');
    }
  }
  templateString = copy;
  //now any conditionals removed we can do simple substitution
  var key, find, re;
  for (key in data) {
    find = '\\$\\{\\s*' + key + '\\s*\\}';
    re = new RegExp(find, 'g');
    templateString = templateString.replace(re, data[key]);
  }
  return templateString;
}
