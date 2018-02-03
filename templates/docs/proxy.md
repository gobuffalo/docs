<%= h1("Using a proxy") %>

Buffalo can be used raw to handle queries. You just need to start your app and listen on a given address and port:

```bash
# Env config
ADDR=0.0.0.0
PORT=80

# Start your app as a daemon, for example:
./myapp &
```

On common scenarios though, you'll be using a proxy to distribute queries to a cluster, handle cases when your app is down, and so on.

<%= title("NGINX") %>

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

### Using an UNIX domain socket

<%= sinceVersion("0.10.3") %>

[UNIX sockets](https://en.wikipedia.org/wiki/Unix_domain_socket) are a common way to do inter-process communication (IPC) on UNIX systems. This means a program **A** can talk to a program **B**, using a file descriptor, just like they do using the TCP stack.

In our case, this allows you to have an instance of Buffalo running behind the proxy, without having to handle the full TCP stack between Buffalo and the proxy. This way, your app will answer faster!

**app env config:**
```bash
ADDR=unix:/tmp/buffalo.sock
```

**NGINX config:**
```nginx
upstream buffalo_app {
    server server unix:/tmp/buffalo.sock;
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

<%= title("Apache 2") %>

### Using an IP address

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