#!/usr/bin/env bash

cat resources | awk -v id="$ID" '{ if ($0 != id){ print($0) }}' > resources
