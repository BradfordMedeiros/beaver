#!/usr/bin/env bash

REPO=$(echo $OPTIONS | awk '{ if ($1 == "repo") { print($2) }}')
echo "$ID $REPO" >> resources 
./start-server.sh