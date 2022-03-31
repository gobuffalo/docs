## Détruire une ressource

Vous pouvez supprimer les fichiers générés par ce générateur en exécutant la commande suivante :

```bash
$ buffalo destroy resource users
```

Cette commande vous demandera de confirmer la suppression de chaque fichier. Vous pouvez soit répondre à chaque fois avec y/n, ou utiliser le flag `-y` pour tout accepter d'un coup :

```bash
$ buffalo destroy resource users -y
```

Ou dans sa forme courte :

```bash
$ buffalo d r users -y
```
