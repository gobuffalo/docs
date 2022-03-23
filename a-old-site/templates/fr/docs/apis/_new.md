## Créer une nouvelle API

Lors de la création d'une nouvelle application Buffalo en utilisant la commande `buffalo new`, le flag optionnel `--api` permet de générer une application qui est optimisée pour servir des points d'API, au lieu de servir une application Buffalo complète.

```bash
$ buffalo new --api coke
```

### Arborescence simplifiée

Les applications générées avec le flag `--api` ne contiennent pas de configuration pour le frontend. Cela signifie qu'il n'y a pas de support pour les templates, les feuilles de style, etc.

<div class="row">
    <div class="col-md-6 col-sm-12">
        <h5><code>buffalo new coke --api</code></h5>
        <%= partial("en/docs/apis/api_ls.md") %>
    </div>
    <div class="col-md-6 col-sm-12">
        <h5><code>buffalo new coke</code></h5>
        <%= partial("en/docs/apis/web_ls.md") %>
    </div>
</div>

### Fichiers `actions/app.go actions/render.go` modifiés

Les applications API ont les fichiers `actions/app.go` et `actions/render.go` modifiés de sorte à servir de base pour une API.

<h5><code>buffalo new coke --api</code></h5>

<%= partial("en/docs/apis/api_app.md") %>
<%= partial("en/docs/apis/api_render.md") %>

<h5><code>buffalo new coke</code></h5>

<%= partial("en/docs/apis/web_app.md") %> 
<%= partial("en/docs/apis/web_render.md") %>