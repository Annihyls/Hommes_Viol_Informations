# Hommes Viol Informations

## Requirement
- Golang 1.27.0 or newer
- Docker 29.8.2 or newer

## Set up

### Build
First you need to create a .env file that contains this variables :
```
# ----- Database user definition ------#
USER_POSTGRES_LOGIN=""
USER_POSTGRES_PASSWORD=""

USER_INTERNAL_LOGIN=""
USER_INTERNAL_PASSWORD=""

USER_EXTERNAL_LOGIN=""
USER_EXTERNAL_PASSWORD=""

HVI_SERVER_CONF=""
# ------------------------------------#
```
You replace "" by your wanted configuration.

Then you can execute export_env.sh
```
./export_env.sh
```

If it a success, you can execute setup_database.sh
```
./setup_database.sh
```

If it also a success, execute the following command to start the server:
```
cd hvi_server/
./start_server.sh
```

And if you want running tests:
```
./start_server.sh test
```
