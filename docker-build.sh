#! /bin/sh
docker build -t airbit.uit.no $(dirname "$(readlink -f "$0")")
