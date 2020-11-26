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

echo "Choose an option:"
echo "1- Run Tests"
echo "2- Run Prod"
echo "3- Run Tests and Prod"
echo "4- Purge All"

read opt

if [ $opt == "1" ]; then
    # Running tests 
    docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
    docker-compose down
elif [ $opt == "2" ]; then
    # Deploying/Up Containeres/Services
    echo "- Building and deploying containeres"
    docker-compose build && docker-compose up
elif [ $opt == "3" ]; then
    # Running tests 
    docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
    docker-compose down
    # Deploying/Up Containeres/Services
    echo "- Building and deploying containeres"
    docker-compose build && docker-compose up
elif [ $opt == "4" ]; then
    docker-compose down --remove-orphans --volumes
    docker system prune -a --volumes
fi

