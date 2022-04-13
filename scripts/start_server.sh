#!/bin/bash
cd /opt/simple-reddit
cp /home/ubuntu/.env /opt/simple-reddit
sudo /usr/local/go/bin/go build -o simple-reddit-build main.go  # due to permission denied error
echo "Running server..."
nohup ./simple-reddit-build &; disown
echo "Server is now up!"
 