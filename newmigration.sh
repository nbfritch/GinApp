#!/usr/bin/bash
TIME=$(date +%s)
set -u
if [ -z $1 ]; then
    echo "Migration name is required"
else
    touch ./migrations/$TIME"_"$1.down.sql
    touch ./migrations/$TIME"_"$1.up.sql
fi