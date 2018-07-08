# bob.expert
A collection of go commands (for now) that produce content for a webserver, along with a caddy configuration.

Future work may split up the caddy config so that this is a pure go repo.

```
alias refresh='sudo systemctl daemon-reload'
alias caddystart='sudo systemctl start caddy'
alias caddystatus='sudo systemctl status caddy'
alias caddystop='sudo systemctl stop caddy'
```

If webserver is down:
```
caddystart
```

To check status
```
caddystatus
```
