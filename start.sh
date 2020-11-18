#!/bin/bash

# Checking if is running as sudo/root
if [ "$EUID" -ne 0 ]
  then echo "Please run as root"
  exit
fi

# Env file
ENV=.env

# Checking if .env file exists
if [ ! -f "$ENV" ]; then
    echo "- $ENV does not exists"
    echo "- Copying .env.example into .env ..."
    cp .env.example .env
    echo "- Copied!"
    echo "- Now you can configure the .env file as you wish"
    exit
fi

# Deploying/Up Containeres/Services
echo "- Building and deploying containeres"
docker-compose build && docker-compose up
