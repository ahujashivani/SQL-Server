#use when smoke-testing main.go
#export necessary ENV vars to run main.go

export DB_PORT="8000"
echo "DB_PORT exported."
export DB_USER="root"
echo "DB_USER exported."
export DB_PASSWORD="rootpassword"
echo "DB_PASSWORD exported."
export DB_HOST="127.0.0.1:3306"
echo "DB_HOST exported."
export DB_DATABASE="Dummy_data"
echo "DB_DATABASE exported."
echo "All env var exported to terminal."
