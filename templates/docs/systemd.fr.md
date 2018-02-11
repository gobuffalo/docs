<%= h1("Systemd service") %>

Dans ce chapitre, nous allons voir comment installer une app Buffalo comme un service Systemd. Systemd est le nouveau standard de beaucoup de distributions GNU/Linux, pour éxecuter les services systèmes.

Il vous permet de configurer votre application d'une manière standard, et de gérer son cylcle de vie avec les commandes `systemctl`.

<%= title("Installez votre app Buffalo") %>

La première étape est de placer votre application dans le bon dossier: sur Debian, la place habituel pour les applications installées à la main est `/usr/local/bin`. C'est là que nous allons installé l'application.

```bash
$ sudo mv myapp /usr/local/bin
```

Assurez vous que les droits sont mis correctements, and donnez les droits à l'utilisateur que vous voulez utilisez. Ici, je vais utilisé le compte `root`.

```bash
$ sudo chown root: /usr/local/bin/myapp
$ sudo chmod +x /user/local/bin/myapp
```

<%= title("Créez un fichier de config systemd") %>

Les fichiers de services systemd sont situés dans `/lib/systemd/system/`, nous allons créer un nouveau fichier `myapp.service` pour notre app.

```ini
[Unit]
Description=My super app

[Service]
ExecStart=/usr/local/bin/myapp
User=root
Group=root
UMask=007

[Install]
WantedBy=multi-user.target
```

Ici nous créons a nouveau service qui a pour nom lisible "My super app". C'est un service simple, ce qui va créer un nouveau processus décrit avec `ExecStart`: le chemin absolu de notre application. The processus va être éxecuter en tant que `root:root`, avec un `UMask` qui va donner les droits seulement au propriétaire du processus(root).

Dans la secition `Install`, nous allons juste dire à Systemd d'attendre que le système soit prêt. Si vous avez d'autres exigences, vous pouvez demandez à Systemd d'attendre pour une base de données, par exemple:

```ini
[Unit]
Description=My super app
After=mysql.service

[Service]
ExecStart=/usr/local/bin/myapp
User=root
Group=root
UMask=007

[Install]
WantedBy=multi-user.target
```

<%= title("Déclarer les variables d'environment") %>

La manière officielle de gérer la configuration avec Buffalo est à travers les [variables d'environment](/docs/env-vars). En utilisant Systemd, vous pouvez les définir avec un fichier de remplacement.

Notre fichier de remplacement est situé dans  `/etc/systemd/system/myapp.service.d/`, et est appelé `override.conf`.

```ini
[Service]
Environment="ADDR=0.0.0.0"
Environment="GO_ENV=production"
Environment="SESSION_SECRET=kqdjmlkajdùméa]$"
```

Chaque ligne `Environment` définit une variable d'environment dans votre app.

<%= title("Jouez avec le service") %>

Le service systemd est maintenant prêt, vous pouvez le tester avec les commandes systemctl and journalctl:

```bash
$ sudo systemctl start myapp.service
```

Pour démarrer le service, and vérifier que tout fonctionne normallement.

```bash
$ journalctl -u myapp.service -f
```

La lecture des logs depuis la sortie standard (-u pour définir le nom du service, -f pour suivre les logs).

```bash
$ sudo systemctl stop myapp.service
```

Pour stoper le service, pour éffectuer une maintenance (par example).

<%= title("Activez le service au démarrage") %>

Une fois que le service fonctionne comme vous les voulez, vous pouvez l'activer au démarrage. De cette manière, si le serveur doit redémarrer, votre app va redémarrer aussi.

```bash
$ sudo systemctl enable myapp.service
```
