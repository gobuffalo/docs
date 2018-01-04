# Using a proxy

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

**app env config**
```bash
ADDR=127.0.0.1
PORT=3000
```

**nginx config:**
```nginx
upstream buffalo_app {
    server 127.0.0.1:3000;
}

server {
    listen 80;
    server_name example.com;

    location / {
        proxy_pass http://buffalo_app;
    }
}
```

#### Multiple backend apps

Using different ports just for the example:

**app1 env config**
```bash
ADDR=0.0.0.0 
PORT=3000
```

**app2 env config**
```bash
ADDR=0.0.0.0 
PORT=3001
```

**app3 env config**
```bash
ADDR=0.0.0.0 
PORT=3002
```

**nginx config:**
```nginx
upstream buffalo_app_hosts {
    server host1.example.com:3000;
    server host2.example.com:3001;
    server host3.example.com:3002;
}

server {
    listen 80;
    server_name example.com;

    location / {
        proxy_pass http://buffalo_app_hosts;
    }
}
```

### Using an UNIX domain socket

<%= sinceVersion("0.10.3") %>

**app env config**
```bash
ADDR=unix:/tmp/buffalo.sock
```

**nginx config:**
```nginx
upstream buffalo_app {
    server server unix:/tmp/buffalo.sock;
}

server {
    listen 80;
    server_name example.com;

    location / {
        proxy_pass http://buffalo_app;
    }
}
```