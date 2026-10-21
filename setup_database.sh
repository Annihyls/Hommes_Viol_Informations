#!/bin/bash

validate_variable() {
	local variable_name="$1"
	local variable_value="$2"
	if [[ -z "$variable_value" ]]; then
		echo "Error: $variable_name must be set in your environment" >&2
		exit 1
	fi
	if [[ "$variable_value" == *" "* ]]; then
        	echo "Error: $variable_name must NOT contains spaces" >&2
        	exit 1
        fi
}

set -e # stop script when an error occurred
validate_variable USER_POSTGRES_LOGIN "$USER_POSTGRES_LOGIN"
validate_variable USER_POSTGRES_PASSWORD "$USER_POSTGRES_PASSWORD"
validate_variable USER_INTERNAL_LOGIN "$USER_INTERNAL_LOGIN"
validate_variable USER_INTERNAL_PASSWORD "$USER_INTERNAL_PASSWORD"
validate_variable USER_EXTERNAL_LOGIN "$USER_EXTERNAL_LOGIN"
validate_variable USER_EXTERNAL_PASSWORD "$USER_EXTERNAL_PASSWORD"

DATABASE_NAME="hvi_database"
DATABASE_PORT="5432"
PATH_TO_PROJECT="$HOME/Projects/homme_viol_infos"

sudo docker rm -f postgres-container
echo "Container erased"

sudo rm -rf "$PATH_TO_PROJECT/hvi_datas/18/"
echo "Datas erased at $PATH_TO_PROJECT/hvi_datas/18/"

sudo --preserve-env=DATABASE_PORT,USER_POSTGRES_LOGIN,USER_POSTGRES_PASSWORD,PATH_TO_PROJECT docker compose up -d
echo "Waiting for PostgreSQL..."

until sudo docker exec postgres-container pg_isready -U postgres >/dev/null 2>&1
do
    sleep 1
done
echo "PostgreSQL is ready"

sudo docker exec -i postgres-container psql -v ON_ERROR_STOP=1 -U $USER_POSTGRES_LOGIN <<EOSQL
	CREATE DATABASE $DATABASE_NAME OWNER $USER_POSTGRES_LOGIN;
EOSQL
echo "Database created"

sudo docker exec -i postgres-container psql -v ON_ERROR_STOP=1 -U $USER_POSTGRES_LOGIN -d "$DATABASE_NAME" <<EOSQL
	CREATE USER $USER_INTERNAL_LOGIN WITH PASSWORD '$USER_INTERNAL_PASSWORD';
	CREATE USER $USER_EXTERNAL_LOGIN WITH PASSWORD '$USER_EXTERNAL_PASSWORD';
	
	GRANT CONNECT ON DATABASE $DATABASE_NAME TO $USER_EXTERNAL_LOGIN;
	GRANT USAGE ON SCHEMA public TO $USER_EXTERNAL_LOGIN;

	GRANT CONNECT ON DATABASE $DATABASE_NAME TO $USER_INTERNAL_LOGIN;
	GRANT USAGE ON SCHEMA public TO $USER_INTERNAL_LOGIN;
EOSQL
echo "Users created"

sudo docker exec -i postgres-container psql -v ON_ERROR_STOP=1 -U $USER_POSTGRES_LOGIN -d $DATABASE_NAME < "$PATH_TO_PROJECT/hvi_sql/schema.sql"
echo "Schemas applied"

sudo docker exec -i postgres-container psql -v ON_ERROR_STOP=1 -U $USER_POSTGRES_LOGIN -d $DATABASE_NAME < "$PATH_TO_PROJECT/hvi_sql/data.sql"
echo "Datas added"

sudo docker exec -i postgres-container psql -v ON_ERROR_STOP=1 -U $USER_POSTGRES_LOGIN -d "$DATABASE_NAME" <<EOSQL
	GRANT SELECT ON ALL TABLES IN SCHEMA public TO $USER_EXTERNAL_LOGIN;
	GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO $USER_INTERNAL_LOGIN;
EOSQL
echo "User privileges updated"
