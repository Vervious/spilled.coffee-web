#!/usr/bin/env bash

# rebuild go files
export GOPATH=/home/ubuntu/bob.expert/cmds
/usr/lib/go-1.10/bin/go install ./...

# restart caddy
sleep 20
sudo /bin/systemctl restart caddy.service