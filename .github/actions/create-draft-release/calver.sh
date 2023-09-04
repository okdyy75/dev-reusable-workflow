#!/bin/bash
set -eu

##############################
## Get next calver
## @link https://github.com/Connehito/gdp/blob/5a5e9ccec14b9d7a0462cfc6d85667add4ba5a46/format.go#L30C1-L47
##############################
today=$(date +'%Y%m%d')
regexp="(.*)([0-9]{8})\.(.+)"

version=${1:-$today}

if [[ $version =~ $regexp ]]; then
    if [ "$today" = "${BASH_REMATCH[2]}" ]; then
        minor=${BASH_REMATCH[3]}
        next=$(($minor + 1))
        echo "${BASH_REMATCH[1]}${today}.${next}"
        exit 0
    fi
    echo "${BASH_REMATCH[1]}${today}.1"
    exit 0
fi
echo "${today}.1"
exit 0
