#!/bin/bash
export PATH=$PATH:/usr/local/go/bin
go version
cd /opt/simple-reddit
go mod download

sudo mv /opt/simple-reddit/scripts/simple-reddit-backend.service /etc/systemd/system
sudo systemctl daemon-reload