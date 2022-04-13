#!/bin/bash
cd /opt/simple-reddit
go build -o simple-reddit-build main.go
nohup ./simple-reddit-build & disown
