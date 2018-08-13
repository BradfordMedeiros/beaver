#!/usr/bin/env bash

OLD_PID=$(cat hooker.PID 2>/dev/null)

if [ ! -z "$OLD_PID" ]; then
	echo "killing server"
	kill "$OLD_PID" || (echo "Could not kill old server" && exit 1)
	rm hooker.PID
	rm hooker-config
fi



