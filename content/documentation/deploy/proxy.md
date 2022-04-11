---
name: Using a Proxy
seoDescription: "Using a Proxy"
seoKeywords: ["buffalo", "go", "golang", "proxy", "nginx", "apache"]
weight: 20
aliases:
  - /docs/deploy/proxy
  - /en/docs/deploy/proxy
---
# Using a Proxy

Buffalo can be used raw to handle queries. You just need to start your app and listen on a given address and port:

```bash
# Env config
ADDR=0.0.0.0
PORT=80

# Start your app as a daemon, for example:
./myapp &
```

On common scenarios though, you'll be using a proxy to distribute queries to a cluster, handle cases when your app is down, and so on.

## NGINX

NGINX allows two ways to be configured with your app:

### Using an IP address

#### Single backend app on same host

**app env config:**
```bash
ADDR=127.0.0.1
PORT=3000
```

**NGINX config:**
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

#### Multiple backend apps

Using different ports just for the example:

**app1 env config:**
```bash
ADDR=0.0.0.0
PORT=3000
```

**app2 env config:**
```bash
ADDR=0.0.0.0
PORT=3001
```

**app3 env config:**
```bash
ADDR=0.0.0.0
PORT=3002
```

**NGINX config:**
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

### Using a UNIX domain socket

{{< since "0.10.3" >}}

[UNIX sockets](https://en.wikipedia.org/wiki/Unix_domain_socket) are a common way to do inter-process communication (IPC) on UNIX systems. This means a program **A** can talk to a program **B**, using a file descriptor, just like they do using the TCP stack.

In our case, this allows you to have an instance of Buffalo running behind the proxy, without having to handle the full TCP stack between Buffalo and the proxy. This way, your app will answer faster!

There are a couple of things to note about UNIX sockets. Since a UNIX socket is a file, UNIX file permissions apply. Therefore, whatever user owns the NGINX processes (typically `nginx`) needs to be able to both read from and write to the socket. **Executing `chmod 777` on the socket file will work, but this is almost always a bad idea!** Since, by default, groups have full read/write permissions on sockets created in Buffalo, a simpler and more secure solution would be to add the NGINX user to the user's group that owns the app. The command to do this would be along the lines of `usermod -aG buffalo nginx`.

Socket files are typically created under the `/tmp` directory as in the example
configuration below. However, in some more recent distributions of Linux,
particularly newer [RedHat family](http://fedoraproject.org/wiki/Features/ServicesPrivateTmp) distros, `/tmp` and `/var/tmp` are namespaced so only the user that creates the file can see that it even exists. On these distributions, you will want to use something along the lines of `unix:/var/sock/buffalo.sock` instead of the example address given below.

**app env config:**
```bash
ADDR=unix:/tmp/buffalo.sock
```

**NGINX config:**
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

### Using an IP address

**app env config:**
```bash
ADDR=127.0.0.1
PORT=3000
```

**Apache 2 config:**
```apache
<VirtualHost *:80>
    ProxyPreserveHost On

    # Proxy requests to Buffalo
    ProxyPass / http://0.0.0.0:3000/
    ProxyPassReverse / http://0.0.0.0:3000/

    ServerName example.com
</VirtualHost>
```
