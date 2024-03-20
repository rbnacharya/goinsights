#!/bin/bash

# Check if psql is installed
if ! [ -x "$(command -v psql)" ]; then
    echo "Error: psql is not installed." >&2
    exit 1
fi

# Install psql if not installed

# Load environment variables from .env file
ENV_FILE="$(dirname "$0")/../.env"
if [[ -f "$ENV_FILE" ]]; then
    source "$ENV_FILE"
    echo "Environment variables loaded from $ENV_FILE"
else
    echo "Error: .env file not found."
    exit 1
fi

# Connect to PostgreSQL database
PG_HOST="$DB_HOST"
PG_PORT="$DB_PORT"
PG_DATABASE="$DB_NAME"
PG_USER="$DB_USER"
PG_PASSWORD="$DB_PASSWORD"

# Loop through all scripts in the "db" folder and execute them
for script in ../../db/*.sql; do
    psql -h "$PG_HOST" -p "$PG_PORT" -U "$PG_USER" -d "$PG_DATABASE" -f "$script"
done
