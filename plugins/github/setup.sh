#/usr/bin/env sh

./hooker &

echo $! > "$1_PID"

