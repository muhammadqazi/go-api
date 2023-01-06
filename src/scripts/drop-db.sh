#!/bin/bash

# Load the environment variables from the .env file
source ../../.env

lsof -n -i$PORT | grep LISTEN | awk '{ print $2 }' | uniq | xargs kill -9

# Connect to the database and terminate the connection
docker exec -it $POSTGRES_CONTAINER_NAME psql -U $POSTGRES_USER $POSTGRES_DB -c "SELECT pg_terminate_backend (pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = '$POSTGRES_DB';"

# Connect to the Postgres container and drop the database
docker exec -it $POSTGRES_CONTAINER_NAME psql -U $POSTGRES_USER -c "DROP DATABASE $POSTGRES_DB;"

# Create a new database with the same name
docker exec -it $POSTGRES_CONTAINER_NAME psql -U $POSTGRES_USER -c "CREATE DATABASE $POSTGRES_DB;"

# Change the current working directory to ../../
cd ../../

make fast