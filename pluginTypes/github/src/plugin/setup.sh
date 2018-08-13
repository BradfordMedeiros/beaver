#!/usr/bin/env bash

./hooker &

echo $! > "$1.PID"

