#!/usr/bin/env bash

ROOT=..

if [ -z "$1" ]
  then
    echo "Укажите название сервиса"
    exit 1
fi



source variables.sh
cd $ROOT/docker
docker compose build
docker compose -p channel_helper up --build --force-recreate --no-deps $1