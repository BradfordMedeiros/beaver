#!/usr/bin/env bash

cat resources | awk '{ printf "/"; printf $1; printf  " "; printf "notify-send "; print $1 }' > hooker-config
