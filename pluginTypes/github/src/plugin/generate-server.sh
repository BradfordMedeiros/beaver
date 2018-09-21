#!/usr/bin/env bash

cat resources | awk '{ printf "/"; printf $1; printf  " "; printf $3;  }'
