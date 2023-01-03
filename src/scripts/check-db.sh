#!/bin/bash
# get username password and db name form .env file not that my is two directory up


container_name=$(
    grep -o 'POSTGRES_CONTAINER_NAME=.*' .env | cut -d '=' -f2
    )
username=$(
    grep -o 'POSTGRES_USER=.*' .env | cut -d '=' -f2
    )
password=$(
    grep -o 'POSTGRES_PASSWORD=.*' .env | cut -d '=' -f2
    )
database=$(
    grep -o 'POSTGRES_DB=.*' .env | cut -d '=' -f2
    )


# Check if the container is running
if docker ps -q --filter name=$container_name
then
  # Try to connect to the database inside the container using psql
  if docker exec -e PGPASSWORD=$password -it $container_name psql -U $username -d $database -c '\l'
  then
    # Connection was successful
    echo "Successfully connected to the database"
  else
    # Connection failed
    echo "Failed to connect to the database"
  fi
fi
