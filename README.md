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

## Setup docuementation
  
Compile the binary from the project root directory with:
```
go build -o cowsayaas ./cmd/cowsayaas
```

On your server you will need the binary, and the /data directory.
```
.
├── cowsayaas
└── data
    ├── cows
    │   ├── bunny.cow
    │   ├── default.cow
    │   └── moose.cow
    ├── docs
    │   └── index.html
    └── homepage
        └── index.html
```

The following example describes one way to deploy Cowsay as a Service where:
- The project lives in ```/srv/cowsayaas/```
- The system user ```cowsayaas``` owns the files and process
- The service is managed by systemd
- Caddy is a reverse proxy



