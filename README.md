# Cowsay as a Service

User curl(1) or similar tools to recieve cowsay messages from a server.

Example:

```bash
$ curl 'http://localhost:8080/plain?text=this+is+also+a+way+to+spend+a+sunday&width=15'
 ________________ 
/ this is also a \
| way to spend a |
\ sunday         /
 ---------------- 
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

For now it takes two parameters:

- text=what the cow will say
- width=maximum number of columns of text in one line

There is preliminary support for templates, such as the default.cow file used by default. It is a cow. There is also preliminary support for modifying eyes and tongue through parameters, and for the different standard modes: borg, dead, greedy, paranoia, stoned, tired, wired, and youthful.

Beyond that, the future plan is to also accept requests in JSON and gRPC format.

It has no OS dependencies and uses no third party libraries.

Make sure you run the main.go from the root directory:

```bash
$ go run ./cmd/cowsayaas
```

This idea came to me while taking a shower, and any code in this repository that can be seen as an original work is licensed under the Unlicense license, which means that you are free to do anything you want with it.
