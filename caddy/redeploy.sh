#!/usr/bin/env bash

# rebuild go files
export GOPATH=/home/ubuntu/bob.expert/cmds
go install ./...

# restart caddy
sleep 20
sudo /bin/systemctl restart caddy.service