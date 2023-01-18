#!/bin/bash
set -e

echo "create second db for testing"
psql -v ON_ERROR_STOP=1 --username $POSTGRES_USER --dbname $POSTGRES_DB <<-EOSQL
    CREATE DATABASE $POSTGRES_DB_TEST;
EOSQL
