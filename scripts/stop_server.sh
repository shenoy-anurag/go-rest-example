#!/bin/bash
SIMPLE_REDDIT_PID=$(ps -A -o pid,cmd|grep simple-reddit-build | grep -v grep |head -n 1 | awk '{print $1}')
echo $SIMPLE_REDDIT_PID
if [ -n $SIMPLE_REDDIT_PID ]; then
    kill -9 $SIMPLE_REDDIT_PID
fi
