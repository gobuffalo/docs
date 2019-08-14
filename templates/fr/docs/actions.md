<% seoDescription("Comment définir et utiliser des contrôleurs d'actions ?") %>
<% seoKeywords(["buffalo", "go", "golang", "actions", "contrôleur", "générateur"]) %>

<%= h1("Contrôleurs d'actions") %>

Dans ce chapitre, vous allez apprendre comment les contrôleurs d'actions fonctionnent ; et comment vous pouvez les générer en utilisant les générateurs fournis.

## Qu'est-ce qu'un contrôleur ?

Un contrôleur est la partie *C* du [pattern MVC](https://fr.wikipedia.org/wiki/Mod%C3%A8le-vue-contr%C3%B4leur). C'est la partie qui gère la logique de l'application à partir d'une décision du routeur, et qui produit la réponse appropriée.

Par exemple, si vous demandez le chemin `/` de ce site Web, le contrôleur en charge de la page d'accueil produira une page HTML pour vous, telle que vous la voyez. Si vous construisez une API REST, le contrôleur récupère ou sauvegarde des données, pour demande (poliment) au moteur de rendu de produire la réponse appropriée.

Dans le cas de Buffalo, on parle généralement d'« actions » pour parler des contrôleurs.

## Définir une action

Les actions de Buffalo (ou contrôleurs) sont des fonctions de type <%= doclink("github.com/gobuffalo/buffalo#Handler") %> :

```go
func Home(c buffalo.Context) error {
	return c.Render(200, r.HTML("home.html"))
}
```

Dans cet exemple, nous avons défini une action « Home » et nous avons demandé au moteur de rendu de produire une page HTML, en utilisant le template « home.html », puis de répondre avec un code HTTP 200.

Chaque action prend en paramètre un `buffalo.Context` : consultez [Context](/fr/docs/context) pour savoir tout ce que vous pouvez faire avec.

<%= partial("fr/docs/generators/action.md") %>

## Prochaines étapes

* [Ressources](/fr/docs/resources) - Définir des packs d'actions de type CRUD.
* [Contexte](/fr/docs/context) - En savoir plus sur l'objet `Context` de Buffalo.