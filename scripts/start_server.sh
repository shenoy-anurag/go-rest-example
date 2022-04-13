#!/bin/bash
cd /opt/simple-reddit
cp ~/.env .
go build -o simple-reddit-build main.go
nohup ./simple-reddit-build & disown
