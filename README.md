# Cowsay as a Service

```text
$ curl -X POST -H "Content-Type: application/json" \
> -d '{
> "text": "Code/Setup documentation",
> "width": 15,
> "eyes": "^^"
> }' 'https://cowsay.4d41.se/api'
 _______________ 
/ Code/Setup    \
\ documentation /
 --------------- 
        \   ^__^
         \  (^^)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

(For user documentation, please see the file data/docs/index.html, or go [here](https://cowsay.4d41.se/docs/))

## Code/Setup documentation
  
Compile the binary from the project root directory with:
```
go build -o cowsayaas ./cmd/cowsayaas
```

The listening port can be chosen either with an environment variable called COWPORT, or through setting the -port flag. The flag overrides the environment variable. If neither is specified, it will listen on port 8080.

```text
$ ./cowsayaas -port 8081
2025/06/14 15:27:25 Starting server version 0.1 on port 8081
```
You can run Cowsay as a Service locally, and connect to localhost, but the idea is that it's a service run over the internets. So, if you have ssh access to a Linux server somewhere, the following example describes one way to deploy Cowsay as a Service where:
- The project lives in ```/srv/cowsayaas/```
- The system user ```cowsayaas``` owns the files and process
- The service is managed by systemd
- Caddy is running (for reverse proxy)

These are the commands/instructions for Debian 12, if needed adjust for the system you are using.

1. Log in to your server with a user who has sudo access.
2. Create a system user called cowsayaas:
```bash
server$ sudo useradd --system --no-create-home --shell /usr/sbin/nologin cowsayaas
```
3. Create a place for the project:
```bash
server$ sudo mkdir -p /srv/cowsayaas
```
4. Put the ```cowsayaas``` binary and the ```data/``` directory in ```/srv/cowsayaas/```. From your local computer you can for example run:
```bash
local$ scp -r cowsayaas data/ <username>@<host>:~/
```
and then on your server, copy those files to ```/srv/cowsayaas```:
```bash
server$ sudo cp -r cowsayaas data/ /srv/cowsayaas/
```
5. Change ownership of the files
```bash
server$ sudo chown -R cowsayaas:cowsayaas /srv/cowsayaas
```
6. Create the file ```/etc/systemd/system/cowsayaas.service``` with the appropriate content:
```bash
server$ sudo tee /etc/systemd/system/myservice.service > /dev/null << 'EOF'
[Unit]
Description=Cowsay as a Service
After=network.target

[Service]
Type=simple
User=cowsayaas
Environment=COWPORT=8080
ExecStart=/srv/cowsayaas/cowsayaas
WorkingDirectory=/srv/cowsayaas
Restart=on-failure

[Install]
WantedBy=multi-user.target
EOF
```

7. Reload systemd, enable the service, and start the service
```bash
server$ sudo systemctl daemon-reload
server$ sudo systemctl enable cowsayaas.service
server$ sudo systemctl start cowsayaas.service
```
8. Check that it's working
```bash
server$ sudo systemctl status cowsayaas.service
server$ journalctl -u cowsayaas -f
```

9. Add to your ```/etc/caddy/Caddyfile```. Replace "cowsay.mydomain.com" with your actual domain name. You may have to configure your DNS.
```text
cowsay.mydomain.com {
        reverse_proxy 127.0.0.1:8080
}
```

And that should pretty much be it.

The code is licensed under the Unlicense, which means that you can do basically anything you want with it, and you don't have to attribute me or anyone else.




