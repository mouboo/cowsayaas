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

## Code/Setup docuementation
  
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
- Caddy is a reverse proxy




