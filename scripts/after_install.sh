#!/bin/bash
export PATH=$PATH:/usr/local/go/bin
go version
cd /opt/simple-reddit
go mod download
