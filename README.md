# Homme Viol Infos

## Requirement
- Golang 1.27.0 or newer
- Docker 29.8.2 or newer

## Set up

### Build
First you need to create a .env file that contains this variables :
```
DATABASE_HOST=""
DATABASE_NAME=""
DATABASE_PORT=""

# ----- Database user definition ------#
USER_POSTGRES_LOGIN=""
USER_POSTGRES_PASSWORD=""

USER_INTERNAL_LOGIN=""
USER_INTERNAL_PASSWORD=""

USER_EXTERNAL_LOGIN=""
USER_EXTERNAL_PASSWORD=""
# ------------------------------------#

# ----- System -----#
PATH_TO_PROJECT=""
```
Replace "" by your configuration. Like for example : 
```
PATH_TO_PROJECT="$HOME/homme_viol_info"
```

Then you can execute export_env.sh
```
./export_env.sh
```

If it a success, you can execute setup_database.sh
```
./setup_database.sh
```

If it also a success, execute :
```
cd hvf_server/src
go mod tidy
go run main.go
```
