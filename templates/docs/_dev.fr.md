<%= title("Lancer votre application en mode développement") %>

<%= note() { %>
Avant de démarrer Buffalo pour la première fois, veuillez consulter la documentation sur les [bases de données](/docs/db) pour paramétrer votre application correctement.
<% } %>

L'un des inconvénients du développement en Go est l'absence de «&nbsp;rechargement&nbsp;» de code. Cela signifie qu'à chaque fois que vous modifiez votre code, **vous devez stopper manuellement** votre application, la recompiler et enfin la redémarrer. Buffalo trouve cela ennuyeux, et veut vous faciliter la vie.

```bash
$ buffalo dev
```

La commande `dev` surveille les fichiers `.go`, `.html` et le dossier des [ressources graphiques](/docs/assets) par défaut. Elle se charge de **recompiler et redémarrer votre application** automatiquement, pour que vous n'ayez pas à vous en soucier.

Lancez la commande `buffalo dev` et affichez la page [localhost:3000/](http://localhost:3000/) pour voir tous vos changements en direct&nbsp;!

<figure>
  <img src="/assets/images/new-coke.png" title="Capture d'écran">
  <figcaption>La toute nouvelle application Coke.</figcaption>
</figure>

#### Lancer le serveur de développement sur un port personnalisé

Il arrive souvent que vous ayez déjà une application utilisant le port 3000. Vous pouvez configurer le port utilisé par le serveur de développement en utilisant la variable d'environnement `PORT`&nbsp;:

```bash
$ PORT=3001 buffalo dev
```

Vous pouvez également consulter le chapitre sur les [variables d'environnement](/docs/env-vars) pour plus d'informations sur la configuration de Buffalo.