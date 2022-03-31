<% seoDescription("Systemd Service") %>
<% seoKeywords(["buffalo", "go", "golang", "service", "systemd", "debian", "redhat", "centos"]) %>

<%= h1("Systemd Service") %>

In this chapter, we'll see how to setup your Buffalo app as a Systemd service. Systemd is the new standard on many GNU/Linux distributions, for running the system services.

It allows you to configure an application in a standard way, and manage its lifecycle with `systemctl` commands. You can refer to the [systemd man page](https://www.freedesktop.org/software/systemd/man/systemd.service.html) for further information.

## Install your Buffalo App

The first step is to place your app into the right folder: on Debian, the common place for executables installed by hand is `/usr/local/bin`. That's where we'll install the app.

```bash
$ sudo mv myapp /usr/local/bin
```

Ensure the rights are correctly set, and give the ownership to the user you want to use. Here, I'll use the `root` account.

```bash
$ sudo chown root: /usr/local/bin/myapp
$ sudo chmod +x /usr/local/bin/myapp
```

## Create the systemd config file

The systemd service files are located in `/lib/systemd/system/`, we'll create a new `myapp.service` file there for your app.

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

Here, we create a new service with an readable name "My super app". It's a simple service, which will spawn a new process described with `ExecStart`: the absolute path to your app. This process will be executed as `root:root`, with a `UMask` giving rights only to the process owner (root).

In the `Install` section, we just tell Systemd to wait for a ready system. If you have more requirements, you can ask Systemd to wait for a database, for instance:

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

## Set env variables

The official way to handle config with Buffalo is through [environment variables](/en/docs/getting-started/config-vars). Using Systemd, you can set them with an override file.

Our override file will be located in `/etc/systemd/system/myapp.service.d/`, and be called `override.conf`.

```ini
[Service]
Environment="ADDR=0.0.0.0"
Environment="GO_ENV=production"
Environment="SESSION_SECRET=kqdjmlkajdùméa]$"
```

Each `Environment` line define an environment variable for your app.

## Play with the service

The systemd service is now ready, you can test it with the `systemctl` and `journalctl` commands:

```bash
$ sudo systemctl start myapp.service
```

To start the service, and check if everything is running fine.

```bash
$ journalctl -u myapp.service -f
```

To read the logs from the standard output (`-u` to set the service name, `-f` to follow the logs).

```bash
$ sudo systemctl stop myapp.service
```

To stop the service, for a maintenance (for instance).

## Enable the service on startup

Once the service is working as you want, you can enable it on startup. This way, if the server need to restart, your app will restart as well.

```bash
$ sudo systemctl enable myapp.service
```
