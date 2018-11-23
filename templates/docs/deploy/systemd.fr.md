<% seoDescription("Service Systemd") %>
<% seoKeywords(["buffalo", "go", "golang", "service", "systemd", "debian", "redhat", "centos"]) %>

<%= h1("Service Systemd") %>

Dans ce chapitre, nous allons voir comment installer une app Buffalo comme un service Systemd. Systemd est le nouveau standard de beaucoup de distributions GNU/Linux, pour exécuter les services systèmes.

Il vous permet de configurer votre application d'une manière standard, et de gérer son cycle de vie avec les commandes `systemctl`. Vous pouvez vous référer à la [page man de systemd (EN)](https://www.freedesktop.org/software/systemd/man/systemd.service.html) pour plus d'informations.

<%= title("Installer votre app Buffalo") %>

La première étape est de placer votre application dans le bon dossier&nbsp;: sur Debian, l'emplacement habituel pour les applications installées à la main est `/usr/local/bin`. C'est donc ici que nous allons installer l'application.

```bash
$ sudo mv myapp /usr/local/bin
```

Assurez-vous que les droits soient correctement configurés, et donnez les droits à l'utilisateur que vous voulez utiliser. Ici, je vais utiliser le compte `root`.

```bash
$ sudo chown root: /usr/local/bin/myapp
$ sudo chmod +x /user/local/bin/myapp
```

<%= title("Créer un fichier de config systemd") %>

Les fichiers de services systemd sont situés dans `/lib/systemd/system/`. Nous allons créer un nouveau fichier `myapp.service` pour notre app.

```ini
[Unit]
Description=Ma super app

[Service]
ExecStart=/usr/local/bin/myapp
User=root
Group=root
UMask=007

[Install]
WantedBy=multi-user.target
```

Ici nous créons un nouveau service qui a pour nom «&nbsp;Ma super app&nbsp;». C'est un service simple, ce qui va créer le nouveau processus décrit avec `ExecStart`&nbsp;: le chemin absolu de notre application. Ce processus va être exécuté en tant que `root:root`, avec un `UMask` qui va donner les droits seulement au propriétaire du processus (root).

Dans la section `Install`, nous allons juste dire à Systemd d'attendre que le système soit prêt. Si vous avez d'autres exigences, vous pouvez demander à Systemd d'attendre qu'une base de données soit prête, par exemple&nbsp;:

```ini
[Unit]
Description=Ma super app
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

La manière officielle de gérer la configuration avec Buffalo est à travers les [variables d'environment](/fr/docs/config-vars). En utilisant Systemd, vous pouvez les définir avec un fichier de surcharge.

Notre fichier de surcharge est situé dans  `/etc/systemd/system/myapp.service.d/`, et se nomme `override.conf`.

```ini
[Service]
Environment="ADDR=0.0.0.0"
Environment="GO_ENV=production"
Environment="SESSION_SECRET=kqdjmlkajdùméa]$"
```

Chaque ligne `Environment` définit une variable d'environment dans votre app.

<%= title("Jouer avec le service") %>

Le service systemd est maintenant prêt, vous pouvez le tester avec les commandes `systemctl` and `journalctl`&nbsp;:

```bash
$ sudo systemctl start myapp.service
```

Pour démarrer le service, et vérifier que tout fonctionne normalement.

```bash
$ journalctl -u myapp.service -f
```

La lecture des logs depuis la sortie standard (`-u` pour définir le nom du service, `-f` pour suivre les logs).

```bash
$ sudo systemctl stop myapp.service
```

Pour stopper le service, pour effectuer une maintenance (par example).

<%= title("Activer le service au démarrage") %>

Une fois que le service fonctionne comme vous le souhaitez, vous pouvez l'activer au démarrage. De cette manière, si le serveur doit redémarrer, votre app va redémarrer aussi.

```bash
$ sudo systemctl enable myapp.service
```
