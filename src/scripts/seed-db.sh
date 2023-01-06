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

# move to scripts directory
cd src/scripts

docker exec -it $container_name psql -U $username -d $database -c "`cat seed-db.sql`"
