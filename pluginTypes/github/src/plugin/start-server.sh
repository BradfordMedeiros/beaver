#i!/usr/bin/env bash

./stop-server.sh
./generate-server.sh > ./hooker-config
./hooker -c ./hooker-config &
HOOKER_PID=$!
echo "$HOOKER_PID" > hooker.PID




