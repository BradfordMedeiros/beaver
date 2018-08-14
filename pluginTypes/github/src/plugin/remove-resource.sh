#!/usr/bin/env bash

NEW_CONFIG=$(cat resources | awk -v id="$ID" '{ if ($1 != id){ print($0) }}')
echo "$NEW_CONFIG" > resources
./start-server.sh
