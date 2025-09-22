#!/bin/bash

# Variables
CONTAINER_NAME=my-postgres
DB_USER=admin
DB_PASSWORD=password
DB_NAME=test_db
NETWORK=pgnetwork
HOST_SCHEMA_FILE=./Scripts/schema.sql
HOST_DATA_FILE=./Scripts/sample_data.sql

# Start Postgres container if not already running
if [ "$(docker ps -q -f name=$CONTAINER_NAME)" == "" ]; then
    echo "Starting Postgres container..."
    docker run --name $CONTAINER_NAME \
        --network $NETWORK \
        -e POSTGRES_USER=$DB_USER \
        -e POSTGRES_PASSWORD=$DB_PASSWORD \
        -e POSTGRES_DB=$DB_NAME \
        -p 5432:5432 \
        -d postgres
else
    echo "Postgres container already running."
fi

# Wait a few seconds for DB to be ready
echo "Waiting for Postgres to initialize..."
sleep 5

# Execute schema.sql
echo "Running schema.sql..."
docker exec -i $CONTAINER_NAME psql -U $DB_USER -d $DB_NAME < $HOST_SCHEMA_FILE

# Execute sample_data.sql
echo "Running sample_data.sql..."
docker exec -i $CONTAINER_NAME psql -U $DB_USER -d $DB_NAME < $HOST_DATA_FILE

echo "Database initialized successfully!"