<% seoDescription("Utilisation d'un mandataire (proxy)") %>
<% seoKeywords(["buffalo", "go", "golang", "proxy", "mandataire", "nginx", "apache"]) %>

<%= h1("Utilisation d'un mandataire (proxy)") %>

Buffalo peut-être utilisé pour écouter les requêtes clientes. Vous devez juste démarrer votre application et écouter sur une adresse et un port précis&nbsp;:

```bash
# Env config
ADDR=0.0.0.0
PORT=80

# Start your app as a daemon, for example:
./myapp &
```

Mais dans la plupart des cas, vous allez utiliser un proxy pour distribuer les requêtes à un cluster, gérer les cas où votre application n'est pas démarrée, etc.

## NGINX

NGINX peut-être utilisé de deux manières avec votre app&nbsp;:

### Utilisation avec une adresse IP

#### Une seule application backend sur le même serveur

**Configuration des variables d'environnement :**
```bash
ADDR=127.0.0.1
PORT=3000
```

**Config NGINX :**
```nginx
upstream buffalo_app {
    server 127.0.0.1:3000;
}

server {
    listen 80;
    server_name example.com;

    # Hide NGINX version (security best practice)
    server_tokens off;

    location / {
        proxy_redirect   off;
        proxy_set_header Host              $http_host;
        proxy_set_header X-Real-IP         $remote_addr;
        proxy_set_header X-Forwarded-For   $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;

        proxy_pass       http://buffalo_app;
    }
}
```

#### Plusieurs applications backend

Utilisation de ports différents, juste pour l'exemple:

**Configuration des variables d'environnement de l'app 1 :**
```bash
ADDR=0.0.0.0
PORT=3000
```

**Configuration des variables d'environnement de l'app 2 :**
```bash
ADDR=0.0.0.0
PORT=3001
```

**Configuration des variables d'environnement de l'app 3 :**
```bash
ADDR=0.0.0.0
PORT=3002
```

**Config NGINX :**
```nginx
upstream buffalo_app_hosts {
    server host1.example.com:3000;
    server host2.example.com:3001;
    server host3.example.com:3002;
}

server {
    listen 80;
    server_name example.com;

    # Hide NGINX version (security best practice)
    server_tokens off;

    location / {
        proxy_redirect   off;
        proxy_set_header Host              $http_host;
        proxy_set_header X-Real-IP         $remote_addr;
        proxy_set_header X-Forwarded-For   $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_pass       http://buffalo_app_hosts;
    }
}
```

### Utilisation des Sockets UNIX

<%= sinceVersion("0.10.3") %>

Les [sockets UNIX](https://fr.wikipedia.org/wiki/Berkeley_sockets#Socket_unix) sont un moyen habituel de faire communiquer des processus entre eux (IPC) sur les systèmes UNIX. Cela veut dire qu'un programme **A** peut parler avec un programme **B** en utilisant un descripteur de fichier, avec la même interface que la stack TCP.

Dans notre cas, cela permet à l'instance de Buffalo de fonctionner derrière un proxy, sans avoir à gérer toute la stack TCP entre Buffalo et le proxy. De cette manière votre application sera plus rapide&nbsp;!

Il a cependant quelques points à noter à propos des sockets UNIX. Puisqu'un socket UNIX est un fichier, les permissions inhérantes aux fichiers UNIX s'appliquent. Cela veut donc dire que l'utilisateur propriétaire des processus NGINX (typiquement `nginx`) doit être en mesure de lire et d'écrire sur le socket. **Appliquer un `chmod 777` sur le fichier socket peut fonctionner, mais c'est en général une très mauvaise idée !** Puisque, par défaut, les groupes ont les pleins accès en lecture et écriture sur les sockets créés par Buffalo, une solution plus simple et sécurisée serait d'ajouter l'utilisateur NGINX au groupe utilisateur à qui appartient l'application. La commande permettant de le faire, ressemble à quelque chose de ce genre&nbsp;: `usermod -aG buffalo nginx`.

Les descripteurs de fichiers socket sont typiquement créés sous le répertoire `/tmp` comme dans la configuration exposée ci-dessous. Néanmoins, sur certaines distributions Linux récentes&nbsp;; en particulier pour distributions de la [famille RedHat (EN)](http://fedoraproject.org/wiki/Features/ServicesPrivateTmp), les répertoires `/tmp` et `/var/tmp` sont cloisonnés, de sorte à ce que seul l'utilisateur qui a créé le fichier soit en mesure d'en voir ne serait-ce que l'existance. Sur ces distributions, vous devrez sans doute changer l'adresse du socket en quelque chose comme `unix:/var/sock/buffalo.sock`, à la place de l'adresse d'exemple donnée ci-dessous.

**Configuration des variables d'environnement :**
```bash
ADDR=unix:/tmp/buffalo.sock
```

**Config NGINX :**
```nginx
upstream buffalo_app {
    server unix:/tmp/buffalo.sock;
}

server {
    listen 80;
    server_name example.com;

    # Hide NGINX version (security best practice)
    server_tokens off;

    location / {
        proxy_redirect   off;
        proxy_set_header Host              $http_host;
        proxy_set_header X-Real-IP         $remote_addr;
        proxy_set_header X-Forwarded-For   $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_pass       http://buffalo_app;
    }
}
```

## Apache 2

### Utilisation avec une adresse IP

**app env config:**
```bash
ADDR=127.0.0.1
PORT=3000
```

**Apache 2 config:**
```apache
&lt;VirtualHost *:80&gt;
    ProxyPreserveHost On

    # Proxy requests to Buffalo
    ProxyPass / http://0.0.0.0:3000/
    ProxyPassReverse / http://0.0.0.0:3000/

    ServerName example.com
&lt;/VirtualHost&gt;
```
